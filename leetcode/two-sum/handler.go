package twosum

import (
	"encoding/json"
	"log"
	"net/http"
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

	responseBody, err := json.Marshal(response)
	if err != nil {
		log.Printf("fail to marshal response, err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(responseBody)
}
