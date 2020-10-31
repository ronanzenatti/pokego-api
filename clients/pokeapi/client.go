package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"pokego-api/config"
)

// GetPokemon - Get one pokemon from PokeAPI by name.
func GetPokemon(search string) (Presenter, int) {
	resp, err := http.Get(config.Config.PokeAPIBaseURL + "pokemon/" + search)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	pokemon := Presenter{}
	json.Unmarshal(body, &pokemon)

	return pokemon, resp.StatusCode
}
