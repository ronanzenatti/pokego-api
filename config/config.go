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
	// PokeAPIBaseURL -  PokeAPI base url.
	PokeAPIBaseURL string
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
	PokeAPIBaseURL, err = viper.Get("POKEAPI_BASE_URL").(string)

	if !err {
		log.Fatal("Environment variables are not completely defined.")
	}
}
