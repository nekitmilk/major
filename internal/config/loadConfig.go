package config

import (
	"major/internal/models"

	"github.com/spf13/viper"
)

func LoadConfig() (config *models.Config, err error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName(".config")
	viper.SetConfigType("json")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	return
}
