package pokeapi

import (
	"github.com/mitchellh/mapstructure"
)

// Presenter - Pokemon struct for visualization.
type Presenter struct {
	Name  string      `json:"name"`
	Image images      `json:"sprites"`
	Types []pokeTypes `json:"types"`
}

type images struct {
	DefaultImage string `json:"front_default"`
}

type pokeTypes struct {
	Slot int      `json:"slot"`
	Type pokeType `json:"type"`
}

type pokeType struct {
	Name string `json:"name"`
}

// ToMap - Convert struct to map.
func (p Presenter) ToMap() map[string]interface{} {
	var outputMap map[string]interface{}
	mapstructure.Decode(p, &outputMap)
	return outputMap
}
