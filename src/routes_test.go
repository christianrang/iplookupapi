package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Tests server start
// func TestAPIServerStart(t *testing.T) {

// 	r := gin.Default()
// 	LoadRoutes(r)
// 	go func() {
// 		if err := r.Run(); err != nil {
// 			log.Printf("err starting the server %+v", err)
// 		}
// 	}()
// }

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

	r := gin.Default()
	LoadRoutes(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/search/ip/8.8.8.8", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
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
