package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/yeom-c/game-server-grpc-go/helper"

	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/account"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeom-c/game-server-grpc-go/db"
	db_session "github.com/yeom-c/game-server-grpc-go/db/redis/session"
	db_common "github.com/yeom-c/game-server-grpc-go/db/sqlc/common"
)

type accountService struct {
	pb.AccountServiceServer
}

func NewAccountService() *accountService {
	return &accountService{}
}

func (s *accountService) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInRes, error) {
	now := time.Now().UTC()
	account, err := db.Store().CommonQueries.GetAccountByUuid(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_AccountErrNotFoundAccount.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	accountUser, err := db.Store().CommonQueries.GetAccountUserByAccountId(ctx, account.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrorWithStack(err_pb.Code_AccountErrNotFoundAccount.String())
		}
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	_, err = db.Store().GameQueries[accountUser.GameDb].GetUserByAccountUserId(ctx, accountUser.ID)
	if err != nil {
		// user 생성
		if err == sql.ErrNoRows {
			_, err = db.Store().TxCreateUser(ctx, nil, accountUser.GameDb, accountUser.ID)
			if err != nil {
				return nil, err
			}

			_, err = db.Store().CommonQueries.UpdateShardingCountByGameDb(ctx, db_common.UpdateShardingCountByGameDbParams{
				GameDb: accountUser.GameDb,
				Count:  1,
			})
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}
		} else {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	_, err = db.Store().BattleQueries.GetUserByAccountUserId(ctx, accountUser.ID)
	if err != nil {
		// battle user 생성
		if err == sql.ErrNoRows {
			_, err = db.Store().BattleQueries.CreateUser(ctx, accountUser.ID)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	_, err = db.Store().CommonQueries.UpdateAccountUserSignedIn(ctx, db_common.UpdateAccountUserSignedInParams{
		ID:         accountUser.ID,
		SignedInAt: now,
	})
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	// 미처리 전투 확인.
	battleResults, err := db.Store().BattleQueries.GetUnconfirmedBattleResults(ctx, accountUser.ID)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	if len(battleResults) > 0 {
		_, _, _, err = db.Store().TxConfirmBattleResults(ctx, nil, nil, accountUser.GameDb, accountUser.ID, battleResults)
		if err != nil {
			return nil, err
		}
	}

	// 매일 업데이트 처리.
	err = db.Store().TxDailyReset(ctx, nil, now, accountUser.GameDb, accountUser.ID)
	if err != nil {
		return nil, err
	}

	mySession := db_session.Session{
		WorldId:       account.WorldID,
		AccountId:     account.ID,
		AccountUserId: accountUser.ID,
		GameDb:        accountUser.GameDb,
		Nickname:      accountUser.Nickname,
		SignedInAt:    now.String(),
	}
	err = db.Store().SessionRedisQueries.CreateAccountSession(ctx, &mySession)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_SessionErr.String())
	}

	return &pb.SignInRes{
		Session: &model_pb.SessionData{
			SessionId:     mySession.SessionId,
			AccountId:     account.ID,
			AccountUserId: accountUser.ID,
			Uuid:          account.Uuid,
			ProfileIdx:    account.ProfileIdx.Int32,
			Nickname:      accountUser.Nickname,
		},
	}, nil
}

func (s *accountService) SignUp(ctx context.Context, req *pb.SignUpReq) (*pb.SignInRes, error) {
	now := time.Now().UTC()

	if req.Uuid == "" || req.Nickname == "" {
		return nil, helper.ErrorWithStack(err_pb.Code_ServerErrInvalidParameter.String())
	}

	account, err := db.Store().CommonQueries.GetAccountByUuid(ctx, req.Uuid)
	accountId := account.ID
	if err != nil {
		// account 생성
		if err == sql.ErrNoRows {
			result, err := db.Store().CommonQueries.CreateAccount(ctx, db_common.CreateAccountParams{
				Uuid:    req.Uuid,
				WorldID: req.WorldId,
				ProfileIdx: sql.NullInt32{
					Int32: req.ProfileIdx,
					Valid: true,
				},
			})
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}

			lastId, err := result.LastInsertId()
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}
			accountId = int32(lastId)
		} else {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	accountUser, err := db.Store().CommonQueries.GetAccountUserByAccountId(ctx, accountId)
	accountUserId := accountUser.ID
	gameDb := accountUser.GameDb
	if err != nil {
		// account user 생성
		if err == sql.ErrNoRows {
			sharding, err := db.Store().CommonQueries.GetGameDb(ctx)
			if err != nil {
				if err != sql.ErrNoRows {
					return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
				}
			}
			gameDb = sharding.GameDb

			existNicknameUser, err := db.Store().CommonQueries.GetAccountUserByNickname(ctx, req.Nickname)
			if err != nil {
				if err != sql.ErrNoRows {
					return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
				}
			}
			if existNicknameUser.Nickname == req.Nickname {
				return nil, helper.ErrorWithStack(err_pb.Code_AccountErrDuplicateNickname.String())
			}

			result, err := db.Store().CommonQueries.CreateAccountUser(ctx, db_common.CreateAccountUserParams{
				AccountID: accountId,
				GameDb:    gameDb,
				Nickname:  req.Nickname,
			})
			if err != nil {
				if err.(*mysql.MySQLError).Number == 1062 {
					return nil, helper.ErrorWithStack(err_pb.Code_AccountErrDuplicateNickname.String())
				}

				return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}

			lastId, err := result.LastInsertId()
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}
			accountUserId = int32(lastId)
		} else {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	_, err = db.Store().GameQueries[gameDb].GetUserByAccountUserId(ctx, accountUserId)
	if err != nil {
		// user 생성
		if err == sql.ErrNoRows {
			_, err = db.Store().TxCreateUser(ctx, nil, gameDb, accountUserId)
			if err != nil {
				return nil, err
			}

			_, err = db.Store().CommonQueries.UpdateShardingCountByGameDb(ctx, db_common.UpdateShardingCountByGameDbParams{
				GameDb: gameDb,
				Count:  1,
			})
			if err != nil {
				return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
			}
		} else {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	_, err = db.Store().BattleQueries.GetUserByAccountUserId(ctx, accountUserId)
	if err != nil {
		// battle user 생성
		if err == sql.ErrNoRows {
			_, err = db.Store().BattleQueries.CreateUser(ctx, accountUserId)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
		}
	}

	mySession := db_session.Session{
		WorldId:       req.WorldId,
		AccountId:     accountId,
		AccountUserId: accountUserId,
		GameDb:        gameDb,
		Nickname:      req.Nickname,
		SignedInAt:    now.String(),
	}
	err = db.Store().SessionRedisQueries.CreateAccountSession(ctx, &mySession)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}

	return &pb.SignInRes{
		Session: &model_pb.SessionData{
			SessionId:     mySession.SessionId,
			AccountId:     accountId,
			AccountUserId: accountUserId,
			Uuid:          req.Uuid,
			Nickname:      req.Nickname,
			ProfileIdx:    req.ProfileIdx,
		},
	}, nil
}
