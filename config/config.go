package config

import (
	"github.com/spf13/viper"
)

// Configurations - APP configurations.
type Configurations struct {
	APPEnv         string `mapstructure:"APP_ENV"`
	APIPort        string `mapstructure:"PORT"`
	PokeAPIBaseURL string `mapstructure:"POKEAPI_BASE_URL"`
}

// Config - Exported configurations.
var Config Configurations

func init() {
	viper.BindEnv("APP_ENV")
	viper.BindEnv("PORT")
	viper.BindEnv("POKEAPI_BASE_URL")

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	viper.Unmarshal(&Config)
}
