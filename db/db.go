package db

import (
	"database/sql"
	"github.com/go-redis/redis/v8"
	"sync"

	"github.com/rs/zerolog/log"
	redis_battle "github.com/yeomc/game-server-grpc-go/db/redis/battle"
	db_battle "github.com/yeomc/game-server-grpc-go/db/sqlc/battle"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yeomc/game-server-grpc-go/config"
	redis_session "github.com/yeomc/game-server-grpc-go/db/redis/session"
	db_common "github.com/yeomc/game-server-grpc-go/db/sqlc/common"
	db_game "github.com/yeomc/game-server-grpc-go/db/sqlc/game"
	db_static_data "github.com/yeomc/game-server-grpc-go/db/sqlc/static_data"
)

var once sync.Once
var instance *store

type store struct {
	SessionRedisDb      *redis.Client
	SessionRedisQueries *redis_session.Queries
	BattleRedis         *redis.Client
	BattleRedisQueries  *redis_battle.Queries
	WorldSessionDb      *redis.Client
	StaticDataDb        *sql.DB
	StaticDataQueries   *db_static_data.Queries
	CommonDb            *sql.DB
	CommonQueries       *db_common.Queries
	GameDb              map[int32]*sql.DB
	GameQueries         map[int32]*db_game.Queries
	BattleDb            *sql.DB
	BattleQueries       *db_battle.Queries
}

func Store() *store {
	once.Do(func() {
		sessionRedis := redis.NewClient(&redis.Options{
			Addr:     config.Config().RedisSessionAddr,
			Username: config.Config().RedisSessionUsername,
			Password: config.Config().RedisSessionPassword,
			DB:       config.Config().RedisSessionDb,
		})

		battleRedis := redis.NewClient(&redis.Options{
			Addr:     config.Config().RedisBattleAddr,
			Username: config.Config().RedisBattleUsername,
			Password: config.Config().RedisBattlePassword,
			DB:       config.Config().RedisBattleDb,
		})

		staticDataDb, err := sql.Open(config.Config().DbStaticDataDriver, config.Config().DbStaticDataSource)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to static_data database")
		}

		commonDb, err := sql.Open(config.Config().DbCommonDriver, config.Config().DbCommonSource)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to common database")
		}

		game0Db, err := sql.Open(config.Config().DbGame0Driver, config.Config().DbGame0Source)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to game_0 database")
		}

		game1Db, err := sql.Open(config.Config().DbGame1Driver, config.Config().DbGame1Source)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to game_1 database")
		}

		battleDb, err := sql.Open(config.Config().DbBattleDriver, config.Config().DbBattleSource)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to battle database")
		}

		instance = &store{
			SessionRedisDb:      sessionRedis,
			SessionRedisQueries: redis_session.New(sessionRedis),
			BattleRedis:         battleRedis,
			BattleRedisQueries:  redis_battle.New(battleRedis),
			StaticDataDb:        staticDataDb,
			StaticDataQueries:   db_static_data.New(staticDataDb),
			CommonDb:            commonDb,
			CommonQueries:       db_common.New(commonDb),
			GameDb:              map[int32]*sql.DB{0: game0Db, 1: game1Db},
			GameQueries: map[int32]*db_game.Queries{
				0: db_game.New(game0Db),
				1: db_game.New(game1Db),
			},
			BattleDb:      battleDb,
			BattleQueries: db_battle.New(battleDb),
		}
	})

	return instance
}
