package config

import (
	"fmt"

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
	fmt.Println(APIEnv)
	// if !err {
	// 	log.Fatal("Environment variables are not completely defined.")
	// }

	APIPort, err = viper.Get("API_PORT").(string)
	fmt.Println(APIPort)
	// if !err {
	// 	log.Fatal("Environment variables are not completely defined.")
	// }

	PokeAPIBaseURL, err = viper.Get("POKEAPI_BASE_URL").(string)
	fmt.Println(PokeAPIBaseURL)
	// if !err {
	// 	log.Fatal("Environment variables are not completely defined.")
	// }

	if err {
	}
}
