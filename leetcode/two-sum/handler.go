package twosum

import (
	"errors"
	"net/http"

	"github.com/kevindoubleu/coding-problem-solver-service/router/middleware"
)

// https://leetcode.com/problems/two-sum/description/

type Input struct {
	Nums   []int `json:"nums"`
	Target int   `json:"target"`
}

type Output struct {
	Input Input `json:"input"`

	Answer []int  `json:"answer,omitempty"`
	Error  string `json:"error,omitempty"`
}

var constraints = map[string]string{
	"nums.length":  "2 <= nums.length <= 10^4",
	"nums[i]":      "-10^9 <= nums[i] <= 10^9",
	"target":       "-10^9 <= target <= 10^9",
	"valid answer": "At least one valid answer exists.",
}

func Handler(w http.ResponseWriter, r *http.Request) {
	input := r.Context().Value("data").(Input)
	output := Output{
		Input: input,
	}

	if err := preCheckConstraints(input.Nums, input.Target); err != nil {
		output.Error = err.Error()
	}

	answer := solve(input.Nums, input.Target)

	if err := postCheckConstraints(answer); err != nil {
		output.Error = err.Error()
	}

	if output.isError() {
		middleware.WriteErrorResponse(output.Error, w)
		return
	}
	output.Answer = answer
	middleware.WriteSuccessResponse(output, w)
}

func preCheckConstraints(nums []int, target int) error {
	length := len(nums)

	if length < 2 || length > 10000 {
		return errors.New("Input not in constraint: " + constraints["nums.length"])
	}

	for i := 0; i < length; i++ {
		if nums[i] < -1000000000 || nums[i] > 1000000000 {
			return errors.New("Input not in constraint: " + constraints["nums[i]"])
		}
	}

	if target < -1000000000 || target > 1000000000 {
		return errors.New("Input not in constraint: " + constraints["target"])
	}

	return nil
}

func postCheckConstraints(answer []int) error {
	if len(answer) == 0 {
		return errors.New("Input not in constraint: " + constraints["valid answer"])
	}

	return nil
}

func (o Output) isError() bool {
	return len(o.Error) > 0
}
