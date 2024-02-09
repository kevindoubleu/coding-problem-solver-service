package twosum

import (
	"net/http"

	"github.com/kevindoubleu/coding-problem-solver-service/router/middleware"
)

// https://leetcode.com/problems/two-sum/description/

type Input struct {
	Nums   []int `json:"nums"`
	Target int   `json:"target"`
}

type Output struct {
	Input  Input `json:"input"`
	Answer []int `json:"answer"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	input := r.Context().Value("data").(Input)

	response := Output{
		Input:  input,
		Answer: solve(input.Nums, input.Target),
	}

	middleware.SuccessResponse(response, w)
}
