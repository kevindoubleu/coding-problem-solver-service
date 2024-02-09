package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func SuccessResponse(data interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Printf("fail to marshal response, err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
