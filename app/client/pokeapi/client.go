package pokeapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetPokemon - Get one pokemon from PokeAPI by name.
func GetPokemon(search string) Presenter {
	resp, err := http.Get(pokeAPIBaseURL + "pokemon/" + search)
	if err != nil {
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	pokemon := Presenter{}
	json.Unmarshal(body, &pokemon)

	return pokemon
}
