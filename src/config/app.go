package config

import (
	"github.com/spf13/viper"
)

type app struct {
	NAME  string
	HOST  string
	PORT  string
	MODE  string
	DEBUG bool
}

func App() app {
	return app{
		NAME:  viper.GetString("APP_NAME"),
		HOST:  viper.GetString("APP_HOST"),
		PORT:  viper.GetString("APP_PORT"),
		MODE:  viper.GetString("APP_MODE"),
		DEBUG: viper.GetBool("APP_DEBUG"),
	}
}
