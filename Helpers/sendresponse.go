package Helpers

import (
	"encoding/json"
	"net/http"

	"github.com/akash-searce/product-catalog/typedefs"
)

func SendResponse(v any, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func SendJResponse(message string, w http.ResponseWriter) {
	SendResponse(typedefs.JResponse{Message: message}, w)
}
