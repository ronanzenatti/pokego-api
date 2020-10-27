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
	viper.AutomaticEnv()
	viper.ReadInConfig()

	setEnvValues()
}

// func setEnvVarFromOS() {
// 	viper.SetDefault("API_ENV", os.Getenv("API_ENV"))
// 	viper.SetDefault("API_PORT", os.Getenv("PORT"))
// 	viper.SetDefault("POKEAPI_BASE_URL", os.Getenv("POKEAPI_BASE_URL"))
// }

func setEnvValues() {
	err := false

	APIEnv, err = viper.Get("API_ENV").(string)
	if !err {
		log.Fatal("Environment variables are not completely defined.")
	}

	APIPort, err = viper.Get("API_PORT").(string)
	if !err {
		log.Fatal("Environment variables are not completely defined.")
	}

	PokeAPIBaseURL, err = viper.Get("POKEAPI_BASE_URL").(string)
	if !err {
		log.Fatal("Environment variables are not completely defined.")
	}
}
