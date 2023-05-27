package core

import (
	"github.com/spf13/viper"
	"gitlab.com/putuyaza/antrian/src/config"
)

func LoadEnv() {
	// set prefix unique environment
	viper.SetEnvPrefix(config.Prefix)
	// set name of environment file name
	viper.SetConfigFile(".env")
	// read environment
	err := viper.ReadInConfig()

	if err != nil {
		panic("Failed to load .env file")
	}
}
