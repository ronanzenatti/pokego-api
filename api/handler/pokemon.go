package handler

import (
	"net/http"

	"pokego-api/api/presenter"
	"pokego-api/clients/pokeapi"
)

var (
	// Pokemon - Handle all function handlers of Pokemon.
	Pokemon = pokemonHandler{}
)

type pokemonHandler struct {
}

func (h pokemonHandler) GetByName(c genContext) error {
	pokemonName := c.Param("pokemon")
	pokeAPI, statusCode := pokeapi.GetPokemon(pokemonName)

	if statusCode == http.StatusNotFound {
		return c.JSON(http.StatusNotFound, presenter.HTTPBody(false, nil))
	}

	pokeAPIMap := pokeAPI.ToMap()

	pokemon := presenter.PokemonPresenter{
		Name:  pokeAPIMap["Name"].(string),
		Image: pokeAPIMap["Image"].(map[string]interface{})["DefaultImage"].(string),
		Types: sliceInterface(pokeAPIMap["Types"].([]pokeapi.PokeTypes)),
	}

	return c.JSON(http.StatusOK, presenter.HTTPBody(true, pokemon))
}
