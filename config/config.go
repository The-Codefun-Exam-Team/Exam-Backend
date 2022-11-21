package config

import (
	"github.com/spf13/viper"
)

type Config struct {

}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Ignore errors when config file is not found
		} else {
			panic("Error while reading config file")
		}
	}

	err = viper.Unmarshal(&config)
	return
}