package config

import (
	"github.com/spf13/viper"
)

// Config is the app configuration structure
type Config struct {
	Database struct {
		Host     string
		Port     string
		User     string
		Name     string
		Password string
	}
}

// Load loads the config file and stores it in the specified variable
func Load(cfg *Config) error {
	// set up config file name
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	// read the file
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// unmarshal the config
	viper.Unmarshal(cfg)
	return nil
}
