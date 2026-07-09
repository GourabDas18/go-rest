package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	var buff bytes.Buffer
	data := map[string]any{
		"success": true,
		"message": "Server is running...",
	}
	err := json.NewEncoder(&buff).Encode(data)

	if err != nil {
		w.Write([]byte("Server is runnng"))
	}

	w.Write(buff.Bytes())
}
