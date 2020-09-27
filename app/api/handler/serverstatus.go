package handler

import (
	"net/http"

	"antares-api/app/api/presenter"
)

var (
	// ServerStatus - Handle all function handlers of ServerStatus.
	ServerStatus = serverStatusHandler{}
)

type serverStatusHandler struct {
}

// GetStatus - Returns the server status.
func (h serverStatusHandler) GetStatus(c genContext) error {
	return c.JSON(http.StatusOK, presenter.HTTPBody(true, map[string]interface{}{
		"serverStatus": "Antares API Online",
	}))
}
