package env

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBAddress string `mapstructure:"DB_ADDRESS"`
	DBDatabaseName string `mapstructure:"DB_DATABASE_NAME"`
	LoggingMode string `mapstructure:"LOGGING_MODE"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Ignore errors when config file is not found
		} else {
			return nil, err
		}
	}

	var config Config
	err = viper.Unmarshal(&config)
	return &config, err
}