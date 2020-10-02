package test

import (
	"net/http"

	"antares-api/api/handler"

	"github.com/labstack/echo"
)

var (
	pokemon      = handler.Pokemon
	serverstatus = handler.ServerStatus
)

type context struct {
	c echo.Context
}

func newContext(req *http.Request, rec http.ResponseWriter) context {
	e := echo.New()

	c := context{
		c: e.NewContext(req, rec),
	}
	return c
}

func (c context) setRouteTester(path string, paramNames []string, paramValues []string) {
	c.c.SetPath(path)

	if paramNames != nil {
		c.c.SetParamNames(paramNames...)
	}

	if paramValues != nil {
		c.c.SetParamValues(paramValues...)
	}

}
