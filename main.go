package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevindoubleu/coding-problem-solver-service/handler"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/ping", handler.PingHandler).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
