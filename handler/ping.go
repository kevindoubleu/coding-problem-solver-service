package handler

import (
	"log"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ping handler called")
	w.Write([]byte("pong"))
}
