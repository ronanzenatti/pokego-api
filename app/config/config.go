package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	// Env - Actual environment (DEV, HOMOL, PROD).
	Env string
	// APIPort - API port.
	APIPort string
)

func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	setEnvValues()
}

func setEnvValues() {
	err := false

	Env, err = viper.Get("ENV").(string)
	APIPort, err = viper.Get("API_PORT").(string)

	if !err {
		log.Fatal("Environment variables are not completely defined.")
	}
}
