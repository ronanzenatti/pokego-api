package presenter

// HTTPBody - Default HTTP body response.
func HTTPBody(success bool, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": success,
		"data":    data,
	}
}
