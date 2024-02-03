package main

import (
	"net/http"

	"github.com/kevindoubleu/coding-problem-solver-service/router"
)

func main() {
	router := router.New()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
