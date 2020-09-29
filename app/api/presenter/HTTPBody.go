package presenter

// DefaultHTTPBody - Default HTTP body.
type DefaultHTTPBody struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// HTTPBody - Default HTTP body response.
func HTTPBody(success bool, data interface{}) DefaultHTTPBody {
	return DefaultHTTPBody{
		Success: success,
		Data:    data,
	}
}
