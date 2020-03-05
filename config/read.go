package config

import (
	"github.com/spf13/viper"
)

func Read(configFile string) error {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.fire")
	}

	return viper.ReadInConfig()
}
