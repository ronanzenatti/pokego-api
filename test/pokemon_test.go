package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemonGetByName(t *testing.T) {
	t.Run("Should return 200 when a valid pokemon name is passed", func(t *testing.T) {
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
		}
	})

	t.Run("Should return 404 when an invalid pokemon name is passed", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := newContext(req, rec)
		c.setRouteTester(
			"/pokeinfo/:pokemon",
			[]string{
				"pokemon",
			},
			[]string{
				"invalid",
			},
		)

		if assert.NoError(t, pokemon.GetByName(c.c)) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
}
