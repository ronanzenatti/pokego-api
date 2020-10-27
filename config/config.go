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
	PORT    int
}

// ServerConfigurations - Server configurations.
type ServerConfigurations struct {
	Port string `json:"port"`
}

// PokeAPIConfigurations - Poke API configurations.
type PokeAPIConfigurations struct {
	BaseURL string `json:"baseURL"`
}

// Config - Exported configurations.
var Config Configurations

func init() {
	viper.SetConfigFile("config.json")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	viper.Unmarshal(&Config)

	fmt.Println(Config)
}
