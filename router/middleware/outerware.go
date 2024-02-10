package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Err string `json:"error"`
}

func WriteErrorResponse(err string, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response, _ := json.Marshal(ErrorResponse{
		Err: err,
	})
	w.Write(response)
}

func WriteSuccessResponse(data interface{}, w http.ResponseWriter) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Printf("fail to marshal response, err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
