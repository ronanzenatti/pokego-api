package presenter

// PokemonPresenter - Pokemon struct for visualization.
type PokemonPresenter struct {
	Name  string      `json:"name"`
	Image images      `json:"images"`
	Types []pokeTypes `json:"types"`
}

type images struct {
	DefaultImage string `json:"default"`
}

type pokeTypes struct {
	Slot int      `json:"slot"`
	Type pokeType `json:"type"`
}

type pokeType struct {
	Name string `json:"name"`
}
