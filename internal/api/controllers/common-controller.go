package controllers

func ResponseError(message string) interface{} {
	result := map[string]interface{}{
		"error":   message,
		"success": false,
	}
	return result
}

func ResponseOK(message string) interface{} {
	result := map[string]interface{}{
		"message": message,
		"success": true,
	}
	return result
}

func Response(success bool, message string) interface{} {
	result := map[string]interface{}{
		"message": message,
		"success": success,
	}
	return result
}
