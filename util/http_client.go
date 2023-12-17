package util

import "net/http"

// HTTPClient is an interface representing an HTTP client
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// RealHTTPClient is an implementation of HTTPClient using the real http.Client
type RealHTTPClient struct{}

// Get performs an HTTP GET request
func (c RealHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
