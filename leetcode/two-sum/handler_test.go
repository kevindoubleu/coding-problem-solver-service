package twosum

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		desc  string
		input Input
		want  Output
	}{
		{
			desc: "happy",
			input: Input{
				Nums:   []int{1, 2, 3},
				Target: 4,
			},
			want: Output{
				Answer: []int{0, 2},
			},
		},
		{
			desc: "constraint violation",
			input: Input{
				Nums:   []int{1, 2, 3},
				Target: 40,
			},
			want: Output{
				Error: "Input not in constraint: At least one valid answer exists.",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			resp := httptest.NewRecorder()
			req := &http.Request{}
			req = req.WithContext(context.WithValue(context.Background(), "data", tC.input))

			Handler(resp, req)
			respBody := Output{}
			json.Unmarshal(resp.Body.Bytes(), &respBody)

			assert.Equal(t, tC.want.Answer, respBody.Answer)
			if tC.want.isError() {
				assert.Equal(t, tC.want.Error, respBody.Error)
			}
		})
	}
}
