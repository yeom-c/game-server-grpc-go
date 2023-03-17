package app

import (
	"context"
	"fmt"
	"github.com/yeom-c/protobuf-grpc-go/gen/golang/protos/error_res"
	"net"
	"runtime/debug"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/yeomc/game-server-grpc-go/config"
	"google.golang.org/grpc"
)

type server struct {
	Grpc *grpc.Server
}

var once sync.Once
var instance *server

func Server() *server {
	once.Do(func() {
		if instance == nil {
			instance = &server{}
			serverOptions := []grpc.ServerOption{
				grpc_middleware.WithUnaryServerChain(
					grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandlerContext(
						func(ctx context.Context, p interface{}) error {
							log.Error().Msgf("%s", p)
							if config.Config().DebugPrintStack {
								debug.PrintStack()
							}
							return status.Error(codes.Internal, error_res.Code_ServerErr.String())
						},
					)),
					logUnaryServerInterceptor(),
					metadataUnaryServerInterceptor(),
					sessionUnaryServerInterceptor(),
				),
				grpc.ConnectionTimeout(3 * time.Minute),
			}

			if config.Config().UseSSL == true {
				creds, _ := credentials.NewServerTLSFromFile("./cert/server-cert.pem", "./cert/server-key.pem")
				serverOptions = append(serverOptions, grpc.Creds(creds))
			}

			instance.Grpc = grpc.NewServer(serverOptions...)
		}

	})

	return instance
}

func (server *server) Run() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Config().ServerPort))
	if err != nil {
		return err
	}
	log.Info().Int("port", config.Config().ServerPort).Msg("started game server")
	return server.Grpc.Serve(listen)
}
