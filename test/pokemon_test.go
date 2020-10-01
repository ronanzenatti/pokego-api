package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonGetByName(t *testing.T) {
	mockJSON := `{"success":true,"data":{"name":"arceus","image":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/493.png","types":[{"slot":1,"type":{"name":"normal"}}]}}`

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := newContext(req, rec)
	c.setRouteTester(
		"/pokeinfo/:pokemon",
		[]string{
			"pokemon",
		},
		[]string{
			"arceus",
		},
	)

	if assert.NoError(t, pokemon.GetByName(c.c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockJSON, strings.TrimSpace(rec.Body.String()))
	}
}
