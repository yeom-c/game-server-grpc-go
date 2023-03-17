package app

import (
	"context"
	"fmt"
	err_pb "github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/yeomc/game-server-grpc-go/config"
	"github.com/yeomc/game-server-grpc-go/db"
	db_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	"github.com/yeomc/game-server-grpc-go/helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var exceptVerificationMethods = []string{
	"/account.AccountService/SignIn",
	"/account.AccountService/SignUp",
}

func logUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		res, err := handler(ctx, req)

		duration := time.Since(startTime)
		statusCode := codes.Unknown
		if st, ok := status.FromError(err); ok {
			statusCode = st.Code()
		}

		md, _ := metadata.FromIncomingContext(ctx)

		logger := log.Info()
		if err != nil {
			logger = log.Error().Err(err)
			if config.Config().DebugPrintStack {
				logger = log.Error().Stack().Err(err)
			}

			err = status.Error(statusCode, err.Error())
		}
		logger.
			Dur("duration", duration).
			Str("status", statusCode.String()).
			Int("status_code", int(statusCode)).
			Str("method", info.FullMethod).
			Interface("metadata", md).
			Interface("request", req.(proto.Message)).
			Interface("response", res.(proto.Message)).
			Msg("")

		return res, err
	}
}

func sessionUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		var session db_session.Session

		// session 확인
		if !helper.Contains(exceptVerificationMethods, info.FullMethod) {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return req.(proto.Message), helper.ErrorWithStack(err_pb.Code_ServerErr.String())
			}

			if len(md["session_id"]) != 1 || len(md["account_id"]) != 1 {
				return req.(proto.Message), helper.ErrorWithStack(err_pb.Code_SessionErrInvalidHeader.String())
			}
			sessionId := md["session_id"][0]
			accountId, _ := strconv.Atoi(md["account_id"][0])

			// session key 비교
			session, _ = db.Store().SessionRedisQueries.GetAccountSession(ctx, int32(accountId))
			if session.SessionId == "" {
				return req.(proto.Message), helper.ErrorWithStack(err_pb.Code_SessionErrEmptySession.String())
			}

			if sessionId != session.SessionId {
				return req.(proto.Message), helper.ErrorWithStack(err_pb.Code_SessionErrInvalidSession.String())
			}

			// set mySession
			ctx = db_session.SetMySession(ctx, session)
		}

		res, err := handler(ctx, req)

		// session 유효시간 연장
		if session.AccountId > 0 {
			sessionKey := fmt.Sprintf("ss_%d", session.AccountId)
			db.Store().SessionRedisQueries.UpdateExpire(ctx, sessionKey, 3600*time.Second)
		}

		return res, err
	}
}

func metadataUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// add metadata
		grpc.SendHeader(ctx, metadata.New(map[string]string{
			"serverTime":      fmt.Sprint(time.Now().Unix()),
			"serverTimeMilli": fmt.Sprint(time.Now().UnixMilli()),
		}))

		return handler(ctx, req)
	}
}
