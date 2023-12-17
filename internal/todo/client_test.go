package todo

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchTodo(t *testing.T) {
	// Positive Test Case
	t.Run("FetchTodo - Successful Fetch", func(t *testing.T) {

		// Create a TodoClient with a mock HTTP client
		client := &TodoClient{client: &MockHTTPClient{}}
		client.client.(*MockHTTPClient).On("Get", API_HOST+"/todos/1").Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"userId": 1, "id": 1, "title": "Test Todo", "completed": false}`)),
		}, nil)

		// Fetch Todo
		result, err := client.FetchTodo(1)
		assert.NoError(t, err)
		assert.Equal(t, Todo{UserID: 1, ID: 1, Title: "Test Todo", Completed: false}, result)
	})

	// Negative Test Cases
	t.Run("FetchTodo - HTTP Error", func(t *testing.T) {
		// Create a TodoClient with a mock HTTP client
		client := &TodoClient{client: &MockHTTPClient{}}
		client.client.(*MockHTTPClient).On("Get", mock.Anything).Return(&http.Response{}, errors.New("HTTP error"))

		// Fetch Todo
		_, err := client.FetchTodo(1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error fetching TODO 1:")
	})

	t.Run("FetchTodo - Decode Error", func(t *testing.T) {
		// Create a TodoClient with a mock HTTP client
		client := &TodoClient{client: &MockHTTPClient{}}
		client.client.(*MockHTTPClient).On("Get", API_HOST+"/todos/1").Return(&http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader("{invalid-json-response")),
		}, nil)

		// Fetch Todo
		_, err := client.FetchTodo(1)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error decoding TODO 1 response:")
	})
}

func TestDecodeResponse(t *testing.T) {
	// Positive Test Case
	t.Run("DecodeResponse - Successful Decode", func(t *testing.T) {
		// Create a response body with valid JSON
		responseBody := `{"userId": 1, "id": 1, "title": "Test Todo", "completed": false}`
		reader := io.NopCloser(strings.NewReader(responseBody))

		// Decode the response
		var result Todo
		err := DecodeResponse(reader, &result)
		assert.NoError(t, err)
		assert.Equal(t, Todo{UserID: 1, ID: 1, Title: "Test Todo", Completed: false}, result)
	})

	// Negative Test Case
	t.Run("DecodeResponse - Decode Error", func(t *testing.T) {
		// Create a response body with invalid JSON
		responseBody := "{invalid-json-response"
		reader := io.NopCloser(strings.NewReader(responseBody))

		// Decode the response
		var result Todo
		err := DecodeResponse(reader, &result)
		assert.Error(t, err)
	})
}
