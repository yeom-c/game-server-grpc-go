package config

import (
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

var cfgOnce sync.Once
var cfgInstance *config

type config struct {
	Env                  string `mapstructure:"ENV"`
	LogLevel             string `mapstructure:"LOG_LEVEL"`
	LogFormat            string `mapstructure:"LOG_FORMAT"`
	DebugPrintStack      bool   `mapstructure:"DEBUG_PRINT_STACK"`
	UseSSL               bool   `mapstructure:"USE_SSL"`
	ServerPort           int    `mapstructure:"SERVER_PORT"`
	RedisSessionAddr     string `mapstructure:"REDIS_SESSION_ADDR"`
	RedisSessionUsername string `mapstructure:"REDIS_SESSION_USERNAME"`
	RedisSessionPassword string `mapstructure:"REDIS_SESSION_PASSWORD"`
	RedisSessionDb       int    `mapstructure:"REDIS_SESSION_DB"`
	RedisBattleAddr      string `mapstructure:"REDIS_BATTLE_ADDR"`
	RedisBattleUsername  string `mapstructure:"REDIS_BATTLE_USERNAME"`
	RedisBattlePassword  string `mapstructure:"REDIS_BATTLE_PASSWORD"`
	RedisBattleDb        int    `mapstructure:"REDIS_BATTLE_DB"`
	DbStaticDataDriver   string `mapstructure:"DB_STATIC_DATA_DRIVER"`
	DbStaticDataSource   string `mapstructure:"DB_STATIC_DATA_SOURCE"`
	DbCommonDriver       string `mapstructure:"DB_COMMON_DRIVER"`
	DbCommonSource       string `mapstructure:"DB_COMMON_SOURCE"`
	DbGame0Driver        string `mapstructure:"DB_GAME_0_DRIVER"`
	DbGame0Source        string `mapstructure:"DB_GAME_0_SOURCE"`
	DbGame1Driver        string `mapstructure:"DB_GAME_1_DRIVER"`
	DbGame1Source        string `mapstructure:"DB_GAME_1_SOURCE"`
	DbBattleDriver       string `mapstructure:"DB_BATTLE_DRIVER"`
	DbBattleSource       string `mapstructure:"DB_BATTLE_SOURCE"`
}

func Config() *config {
	cfgOnce.Do(func() {
		if cfgInstance == nil {
			viper.SetConfigFile(".env")
			viper.AutomaticEnv()

			err := viper.ReadInConfig()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to load config")
			}

			err = viper.Unmarshal(&cfgInstance)
			if err != nil {
				log.Fatal().Err(err).Msg("failed to load config")
			}
		}
	})

	return cfgInstance
}
