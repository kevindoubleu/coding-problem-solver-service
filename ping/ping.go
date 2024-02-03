package ping

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("ping handler called")
	w.Write([]byte("pong"))
}
