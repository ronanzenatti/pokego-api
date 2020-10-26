package config

import (
	"fmt"
	"os"
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
	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()

	setEnvValues()
}

func setEnvValues() {
	// err := false

	// APIEnv, err = viper.Get("API_ENV").(string)
	// log.Fatalln(APIEnv)
	// // if !err {
	// // 	log.Fatal("Environment variables are not completely defined.")
	// // }

	// APIPort, err = viper.Get("API_PORT").(string)
	// fmt.Println(APIPort)
	// // if !err {
	// // 	log.Fatal("Environment variables are not completely defined.")
	// // }

	// PokeAPIBaseURL, err = viper.Get("POKEAPI_BASE_URL").(string)
	// log.Fatalln(PokeAPIBaseURL)
	// // if !err {
	// // 	log.Fatal("Environment variables are not completely defined.")
	// // }

	// if err {
	// }

	fmt.Println(os.Getenv("API_ENV"))
	fmt.Println(os.Getenv("API_PORT"))
	fmt.Println(os.Getenv("POKEAPI_BASE_URL"))
}
