package handler

import (
	"log"
	"net/http"
)

func PingHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("ping handler called")
		w.Write([]byte("pong"))
	})
}
