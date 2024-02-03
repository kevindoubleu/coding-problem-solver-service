package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	Handler(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "pong", resp.Body.String())
}
