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

func TestRoutes(t *testing.T) {

	tests := []struct {
		title      string
		routeTitle string
		route      string
		opt        string
		expected   int
	}{
		{
			"Invalid IP",
			"IP Route",
			"/search/ip",
			"0000.00.00.0",
			406,
		},
		{
			"Valid IP",
			"IP Route",
			"/search/ip",
			"8.8.8.8",
			200,
		},
		{
			"Invalid Domain",
			"Domain Route",
			"/search/domain",
			"a&*.com",
			406,
		},
		{
			"Valid Domain",
			"Domain Route",
			"/search/domain",
			"google.com",
			200,
		},
		{
			"Invalid File Hash",
			"File Hash Route",
			"/search/file_hash",
			"asdfz",
			406,
		},
		{
			"Valid File Hash",
			"File Hash Route",
			"/search/file_hash",
			"74768564ea2ac673e57e937f80c895c81d015e99a72544efa5a679d729c46d5f",
			200,
		},
	}

	r := gin.Default()
	LoadRoutes(r)

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s: %s", test.routeTitle, test.title), func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", test.route, test.opt), nil)
			r.ServeHTTP(w, req)

			// Handles "too many requests" response from VIrusTotal
			if w.Code == 204 {
				// A response code of 204 means the API responded properly, but was timed out by an upstream provider.
				// Skipping the test in this case is to prevent waiting a minute per 4 API requests while signalling to the tester that
				// the test did not execute
				t.SkipNow()
			}

			assert.Equal(t, test.expected, w.Code)
		})
	}

}

// func TestSearchIPRoute(t *testing.T) {

// 	tests := []struct {
// 		title    string
// 		opt      string
// 		expected int
// 	}{
// 		{
// 			"Invalid IP",
// 			"0000.00.00.0",
// 			406,
// 		},
// 		{
// 			"Valid IP",
// 			"8.8.8.8",
// 			200,
// 		},
// 	}

// 	r := gin.Default()
// 	LoadRoutes(r)

// 	for _, test := range tests {
// 		t.Run(test.title, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest("GET", fmt.Sprintf("/search/ip/%s", test.opt), nil)
// 			r.ServeHTTP(w, req)

// 			// Handles "too many requests" response from VIrusTotal
// 			if w.Code == 204 {
// 				time.Sleep(60)
// 				req, _ := http.NewRequest("GET", fmt.Sprintf("/search/ip/%s", test.opt), nil)
// 				r.ServeHTTP(w, req)
// 			}

// 			assert.Equal(t, test.expected, w.Code)
// 		})
// 	}
// }

// func TestSearchDomainRoute(t *testing.T) {

// 	tests := []struct {
// 		title    string
// 		opt      string
// 		expected int
// 	}{
// 		{
// 			"Invalid Domain",
// 			"a&*.com",
// 			406,
// 		},
// 		{
// 			"Valid Domain",
// 			"google.com",
// 			200,
// 		},
// 	}

// 	r := gin.Default()
// 	LoadRoutes(r)

// 	for _, test := range tests {
// 		t.Run(test.title, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest("GET", fmt.Sprintf("/search/domain/%s", test.opt), nil)
// 			r.ServeHTTP(w, req)

// 			// Handles "too many requests" response from VIrusTotal
// 			if w.Code == 204 {
// 				time.Sleep(60)
// 				req, _ := http.NewRequest("GET", fmt.Sprintf("/search/domain/%s", test.opt), nil)
// 				r.ServeHTTP(w, req)
// 			}

// 			assert.Equal(t, test.expected, w.Code)
// 		})
// 	}

// }
// func TestSearchFileHashRoute(t *testing.T) {

// 	tests := []struct {
// 		title    string
// 		opt      string
// 		expected int
// 	}{
// 		{
// 			"Invalid File Hash",
// 			"asdfz",
// 			406,
// 		},
// 		{
// 			"Valid File Hash",
// 			"74768564ea2ac673e57e937f80c895c81d015e99a72544efa5a679d729c46d5f",
// 			200,
// 		},
// 	}

// 	r := gin.Default()
// 	LoadRoutes(r)

// 	for _, test := range tests {
// 		t.Run(test.title, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest("GET", fmt.Sprintf("/search/file_hash/%s", test.opt), nil)
// 			r.ServeHTTP(w, req)

// 			// Handles "too many requests" response from VIrusTotal
// 			if w.Code == 204 {
// 				time.Sleep(60)
// 				req, _ = http.NewRequest("GET", fmt.Sprintf("/search/file_hash/%s", test.opt), nil)
// 				r.ServeHTTP(w, req)
// 			}

// 			assert.Equal(t, test.expected, w.Code)
// 		})
// 	}
// }
