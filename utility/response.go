package utility

import (
	"encoding/json"
	"net/http"
)

type ResponseType string

const (
	Success ResponseType = "Success"
	Error   ResponseType = "Error"
)

func Response(w http.ResponseWriter, status int, message string, data any, responseType ResponseType) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	resp := make(map[string]any)
	if responseType == Success {
		resp["status"] = responseType
		resp["success"] = true
	} else {
		resp["status"] = responseType
		resp["success"] = false
	}

	resp["statusCode"] = status
	resp["message"] = message
	resp["data"] = data
	json.NewEncoder(w).Encode(resp)
}
