package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error      string `json:"error"`
	httpStatus int
}

func NewErrorResponse(err error, httpStatus int) ErrorResponse {
	return ErrorResponse{
		Error:      err.Error(),
		httpStatus: httpStatus,
	}
}

func WriteErrorResponse(data ErrorResponse, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(data.httpStatus)

	response, _ := json.Marshal(data)
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
