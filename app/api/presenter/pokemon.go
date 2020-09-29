package presenter

// PokemonPresenter - Pokemon struct for visualization.
type PokemonPresenter struct {
	Name  string        `json:"name"`
	Image string        `json:"image"`
	Types []interface{} `json:"types"`
}
