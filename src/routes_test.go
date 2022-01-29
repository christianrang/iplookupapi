package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {

	r := gin.Default()
	LoadRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Pong!\"}", w.Body.String())
}

func TestSearchIPRoute(t *testing.T) {

	tests := []struct {
		title    string
		opt      string
		expected int
	}{
		{
			"Invalid IP",
			"0000.00.00.0",
			406,
		},
		{
			"Valid IP",
			"8.8.8.8",
			200,
		},
	}

	r := gin.Default()
	LoadRoutes(r)

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/search/ip/%s", test.opt), nil)
			r.ServeHTTP(w, req)

			assert.Equal(t, test.expected, w.Code)
		})
	}
}

func TestSearchDomainRoute(t *testing.T) {

	r := gin.Default()
	LoadRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search/domain/google.com", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
func TestSearchFileHashRoute(t *testing.T) {

	r := gin.Default()
	LoadRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search/file_hash/74768564ea2ac673e57e937f80c895c81d015e99a72544efa5a679d729c46d5f", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
