package config

import "github.com/spf13/viper"

type database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	Engine   string
}

func Database() database {
	return database{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		Name:     viper.GetString("DB_NAME"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Engine:   viper.GetString("DB_ENGINE"),
	}
}
