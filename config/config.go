package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations - APP configurations.
type Configurations struct {
	Env     string                `json:"env"`
	Server  ServerConfigurations  `json:"server"`
	PokeAPI PokeAPIConfigurations `json:"pokeAPI"`
	PORT    string
	BATATA  string `mapstructure:"PORT"`
}

// ServerConfigurations - Server configurations.
type ServerConfigurations struct {
	Port string `mapstructure:"PORT"`
}

// PokeAPIConfigurations - Poke API configurations.
type PokeAPIConfigurations struct {
	BaseURL string `json:"baseURL"`
}

// Config - Exported configurations.
var Config Configurations

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	viper.BindEnv("PORT")

	viper.Unmarshal(&Config)

	fmt.Println(Config)
}
