package handler

import (
	"net/http"

	"antares-api/app/api/presenter"
	"antares-api/app/client/pokeapi"

	"github.com/mitchellh/mapstructure"
)

var (
	// Pokemon - Handle all function handlers of Pokemon.
	Pokemon = pokemonHandler{}
)

type pokemonHandler struct {
}

func (h pokemonHandler) GetByName(c genContext) error {
	pokemonName := c.Param("pokemon")
	pokeapi := pokeapi.GetPokemon(pokemonName).ToMap()
	pokemon := presenter.PokemonPresenter{}

	mapstructure.Decode(pokeapi, &pokemon)

	return c.JSON(http.StatusOK, presenter.HTTPBody(true, pokemon))
}
