package router

import (
	"github.com/gorilla/mux"
	twosum "github.com/kevindoubleu/coding-problem-solver-service/leetcode/two-sum"
	"github.com/kevindoubleu/coding-problem-solver-service/ping"
	"github.com/kevindoubleu/coding-problem-solver-service/router/middleware"
)

func New() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/ping", ping.Handler).Methods("GET")

	solveRouter := router.PathPrefix("/solve").Subrouter()
	registerLeetcodeSubrouter(solveRouter)

	return router
}

func registerLeetcodeSubrouter(super *mux.Router) {
	subrouter := super.PathPrefix("/leetcode").Subrouter()

	twosumSubrouter := subrouter.PathPrefix("/two-sum").Subrouter()
	twoSumMw := middleware.NewRequestBodyValidatorMiddleware[twosum.Request](true)
	twosumSubrouter.Use(twoSumMw.Middleware)
	twosumSubrouter.HandleFunc("", twosum.Handler).Methods("POST")
}
