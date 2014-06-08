package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHTML(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(HomeHandler))
	defer server.Close()

	resp, err := http.DefaultClient.Get(server.URL)
	assert.Nil(t, err)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/html; charset=utf-8", resp.Header.Get("Content-Type"))
}

func TestHomeJSON(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(HomeHandler))
	defer server.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", server.URL, nil)
	assert.Nil(t, err)
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	assert.Nil(t, err)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}
