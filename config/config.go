package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	// APIEnv - Actual environment (DEV, HOMOL, PROD).
	APIEnv string
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

	APIEnv, err = viper.Get("API_ENV").(string)
	APIPort, err = viper.Get("API_PORT").(string)

	if !err {
		log.Fatal("Environment variables are not completely defined.")
	}
}
