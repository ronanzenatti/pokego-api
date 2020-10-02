package handler

import (
	"net/http"

	"antares-api/api/presenter"
)

var (
	// ServerStatus - Handle all function handlers of ServerStatus.
	ServerStatus = serverStatusHandler{}
)

type serverStatusHandler struct {
}

// GetStatus - Returns the server status.
func (h serverStatusHandler) GetStatus(c genContext) error {
	serverStatus := presenter.ServerStatusPresenter{
		ServerStatus: "Antares API Online",
	}

	return c.JSON(http.StatusOK, presenter.HTTPBody(true, serverStatus))
}
