package service

import (
	"context"
	pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/character"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	model_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/model"

	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	"github.com/yeomc/game-server-grpc-go/helper"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type characterService struct {
	pb.CharacterServiceServer
}

func NewCharacterService() *characterService {
	return &characterService{}
}

func (s *characterService) GetCharacters(ctx context.Context, _ *model_pb.Empty) (*pb.GetCharactersRes, error) {
	var charactersData []*model_pb.CharacterData
	mySession := db_session.GetMySession(ctx)

	characters, err := db.Store().GameQueries[mySession.GameDb].GetCharacterListByAccountUserId(ctx, mySession.AccountUserId)
	if err != nil {
		return nil, helper.ErrorWithStack(err_pb.Code_DatabaseErr.String())
	}
	for _, character := range characters {
		charactersData = append(charactersData, &model_pb.CharacterData{
			Id:             character.ID,
			EnumId:         character.EnumID,
			Exp:            character.Exp,
			EquipmentLevel: character.EquipmentLevel,
			CreatedAt:      timestamppb.New(character.CreatedAt),
		})
	}

	return &pb.GetCharactersRes{
		Characters: charactersData,
	}, nil
}

func (s *characterService) LevelUpSignatureWeapon(ctx context.Context, req *pb.LevelUpSignatureWeaponReq) (*model_pb.Result, error) {
	mySession := db_session.GetMySession(ctx)

	err := db.Store().TxLevelUpSignatureWeapon(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.TargetCharacterId, req.MaterialCharactersId)
	if err != nil {
		return nil, err
	}

	return &model_pb.Result{Result: 1}, nil
}

func (s *characterService) ExtinctCharacter(ctx context.Context, req *pb.ExtinctCharacterReq) (*pb.ExtinctCharacterRes, error) {
	mySession := db_session.GetMySession(ctx)

	reward, err := db.Store().TxExtinctCharacter(ctx, nil, mySession.GameDb, mySession.AccountUserId, req.CharactersId)
	if err != nil {
		return nil, err
	}

	return &pb.ExtinctCharacterRes{
		Reward: &reward,
	}, nil
}
