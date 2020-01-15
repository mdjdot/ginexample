package main

import "testing"

import "net/http/httptest"

import "net/http"

import "github.com/stretchr/testify/assert"

func TestPingRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
