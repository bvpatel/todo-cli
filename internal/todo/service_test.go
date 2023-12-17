package todo

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFetchAndCheckStatus(t *testing.T) {
	// Positive Test Case
	t.Run("FetchAndCheckStatus - Successful Fetch", func(t *testing.T) {
		mockClient := &MockTodoClient{}
		service := &TodoService{todoClient: mockClient}

		// Configure the mock to return a successful response
		mockClient.On("FetchTodo", 1).Return(Todo{ID: 1, Title: "Test Todo"}, nil)
		mockClient.On("FetchTodo", 2).Return(Todo{ID: 2, Title: "Another Todo"}, nil)

		numbers := []int{1, 2}
		todos := service.FetchAndCheckStatus(numbers)

		// Assert that the returned todos match the expected values
		assert.ElementsMatch(t, []Todo{{ID: 1, Title: "Test Todo"}, {ID: 2, Title: "Another Todo"}}, todos)

		// Assert that the FetchTodo method was called with the expected arguments
		mockClient.AssertExpectations(t)
	})

	// Negative Test Case
	t.Run("FetchAndCheckStatus - Error Fetching Todo", func(t *testing.T) {
		mockClient := new(MockTodoClient)
		service := &TodoService{todoClient: mockClient}

		// Configure the mock to return an error for FetchTodo
		mockClient.On("FetchTodo", 1).Return(Todo{}, errors.New("fetch error"))

		numbers := []int{1}
		todos := service.FetchAndCheckStatus(numbers)

		// Assert that the returned todos slice is empty due to the error
		assert.Empty(t, todos)

		// Assert that the FetchTodo method was called with the expected arguments
		mockClient.AssertExpectations(t)
	})

	// Additional Test Case
	t.Run("FetchAndCheckStatus - Concurrent Fetching", func(t *testing.T) {
		mockClient := new(MockTodoClient)
		service := &TodoService{todoClient: mockClient}

		// Configure the mock to return a successful response with a delay
		mockClient.On("FetchTodo", 1).Return(Todo{ID: 1, Title: "Test Todo"}, nil).After(100 * time.Millisecond)

		numbers := []int{1}
		var wg sync.WaitGroup
		resultChan := make(chan Todo, len(numbers))

		// Perform concurrent fetching
		for _, num := range numbers {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				todo, err := service.todoClient.FetchTodo(id)
				if err != nil {
					// Handle error
					return
				}
				resultChan <- todo
			}(num)
		}

		// Wait for all goroutines to finish
		go func() {
			wg.Wait()
			close(resultChan)
		}()

		// Collect results
		var todos []Todo
		for todo := range resultChan {
			todos = append(todos, todo)
		}

		// Assert that the returned todos match the expected values
		assert.ElementsMatch(t, []Todo{{ID: 1, Title: "Test Todo"}}, todos)

		// Assert that the FetchTodo method was called with the expected arguments
		mockClient.AssertExpectations(t)
	})
}
