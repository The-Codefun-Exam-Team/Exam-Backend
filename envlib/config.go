package envlib

import (
	"github.com/spf13/viper"
)

// Config is a struct containing configuration values for various purposes.
type Config struct {
	DBUsername     string `mapstructure:"DB_USERNAME"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBAddress      string `mapstructure:"DB_ADDRESS"`
	DBDatabaseName string `mapstructure:"DB_DATABASE_NAME"`
	LoggingMode    string `mapstructure:"LOGGING_MODE"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
}

// LoadConfig loads config file from current working directorys.
func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")

	// Set config file name to be config.env
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	// Read config file and check for errors
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Ignore errors when config file is not found
		} else {
			return nil, err
		}
	}

	// Unmarshal data to a Config struct
	var config Config
	err = viper.Unmarshal(&config)
	return &config, err
}
