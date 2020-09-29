package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	handler  = pokemonHandler{}
	mockJSON = `{"data":{"name":"arceus","images":{"default":"https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/493.png"},"types":[{"slot":1,"type":{"name":"normal"}}]},"success":true}`
)

func TestGetByName(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/pokeinfo/:pokemon")
	c.SetParamNames("pokemon")
	c.SetParamValues("arceus")

	if assert.NoError(t, handler.GetByName(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockJSON, strings.TrimSpace(rec.Body.String()))
	}
}
