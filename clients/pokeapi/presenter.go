package pokeapi

import "antares-api/pkg/convert"

// Presenter - Pokemon struct for visualization.
type Presenter struct {
	Name  string      `json:"name"`
	Image Images      `json:"sprites"`
	Types []PokeTypes `json:"types"`
}

// Images -  Pokemon images.
type Images struct {
	DefaultImage string `json:"front_default"`
}

// PokeTypes -  Pokemon types list.
type PokeTypes struct {
	Slot int      `json:"slot"`
	Type PokeType `json:"type"`
}

// PokeType -  Pokemon type.
type PokeType struct {
	Name string `json:"name"`
}

// ToMap - Convert struct to map.
func (p Presenter) ToMap() map[string]interface{} {
	var outputMap map[string]interface{}
	convert.MapStruct(p, &outputMap)
	return outputMap
}
