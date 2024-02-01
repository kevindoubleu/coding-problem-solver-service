package main

import (
	"net/http"

	"github.com/kevindoubleu/coding-problem-solver-service/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/ping", handler.PingHandler())

	http.ListenAndServe(":8080", mux)
}
