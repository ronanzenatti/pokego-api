package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerStatusGetStatus(t *testing.T) {
	mockJSON := `{"success":true,"data":{"serverStatus":"Antares API Online"}}`

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := newContext(req, rec)

	c.setRouteTester("/server/status", nil, nil)

	if assert.NoError(t, serverstatus.GetStatus(c.c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, mockJSON, strings.TrimSpace(rec.Body.String()))
	}
}
