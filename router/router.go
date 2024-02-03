package router

import (
	"github.com/gorilla/mux"
	twosum "github.com/kevindoubleu/coding-problem-solver-service/leetcode/two-sum"
	"github.com/kevindoubleu/coding-problem-solver-service/ping"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Handler).Methods("GET")

	solveRouter := router.PathPrefix("/solve").Subrouter()
	leetcodeRouter := solveRouter.PathPrefix("/leetcode").Subrouter()
	leetcodeRouter.HandleFunc("/two-sum", twosum.Handler).Methods("POST")

	return router
}