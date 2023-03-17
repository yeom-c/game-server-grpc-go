package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/yeomc/game-server-grpc-go/app"
	"github.com/yeomc/game-server-grpc-go/config"
	"github.com/yeomc/game-server-grpc-go/service"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	loc, _ := time.LoadLocation("UTC")
	time.Local = loc

	// zerolog Stack() 사용 위함.
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if config.Config().LogFormat != "json" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	logLevel := zerolog.DebugLevel
	switch config.Config().LogLevel {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "fatal":
		logLevel = zerolog.FatalLevel
	case "panic":
		logLevel = zerolog.PanicLevel
	case "disabled":
		logLevel = zerolog.Disabled
	}
	zerolog.SetGlobalLevel(logLevel)

	// server instance 생성, 미들웨어, 서비스, 스케줄러 등록.
	server := app.Server()
	service.NewService(server.Grpc)

	err := server.Run()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
