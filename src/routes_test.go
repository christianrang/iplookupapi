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
	req, _ := http.NewRequest("GET", "/ip/8.8.8.8", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
