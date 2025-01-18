package gourl

import (
	"log"
	"net/http"
	"net/url"
)

// Check if url is valid.
func IsValidUrl(uristr string) bool {
	_, err := url.ParseRequestURI(uristr)
	if err != nil {
		return false
	}

	u, err := url.Parse(uristr)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

// Check if url is reachable.
func TestUrl(uri string) *http.Response {
	// Get request to respective url.
	resp, err := http.Get(uri)

	// Check if error occured.
	if err != nil {
		log.Fatal(err)
	}
	// Return full response to use further on.
	return resp
}
