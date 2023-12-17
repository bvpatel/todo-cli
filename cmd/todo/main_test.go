package todo

import (
	"errors"
	"testing"
	"todo-cli/internal/todo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTodoApp_ReadInputFromFile(t *testing.T) {
	t.Run("ReadInputFromFile - Successful", func(t *testing.T) {
		mockUtil := &MockUtil{}
		app := &TodoApp{inputFile: "input.txt"}
		mockUtil.On("ReadFromFile", "input.txt").Return([]byte("2 4 6 8 10"), nil)
		ReadFromFile = mockUtil.ReadFromFile

		data := app.ReadInputFromFile()
		assert.Equal(t, []byte("2 4 6 8 10"), data)
		mockUtil.AssertExpectations(t)
	})
}

func TestTodoApp_ParseNumbers(t *testing.T) {
	t.Run("ParseNumbers - Successful", func(t *testing.T) {
		app := &TodoApp{}
		numbers := app.ParseNumbers([]byte("2 4 6 8 10"))
		assert.Equal(t, []int{2, 4, 6, 8, 10}, numbers)
	})
}

func TestTodoApp_FetchAndDisplayTodos(t *testing.T) {
	t.Run("FetchAndDisplayTodos - Write to CSV", func(t *testing.T) {
		mockUtil := &MockUtil{}
		mockTodoService := &MockTodoService{}

		app := &TodoApp{
			service:    mockTodoService,
			outputFile: "output.csv",
		}

		mockTodoService.On("FetchAndCheckStatus", []int{2, 4, 6, 8, 10}).Return([]todo.Todo{
			{Title: "Todo1", Completed: false},
			{Title: "Todo2", Completed: true},
		})

		mockUtil.On("WriteToCSV", mock.Anything, "output.csv", []string{"Title", "Completed"}).
			Return(nil)

		WriteToCSV = mockUtil.WriteToCSV
		app.FetchAndDisplayTodos([]byte("2 4 6 8 10"))
		mockUtil.AssertExpectations(t)
		mockTodoService.AssertExpectations(t)
	})

	t.Run("FetchAndDisplayTodos - Write to Console", func(t *testing.T) {
		mockUtil := &MockUtil{}
		mockTodoService := &MockTodoService{}

		app := &TodoApp{
			service: mockTodoService,
		}

		mockTodoService.On("FetchAndCheckStatus", []int{2, 4, 6, 8, 10}).Return([]todo.Todo{
			{Title: "Todo1", Completed: false},
			{Title: "Todo2", Completed: true},
		})

		mockUtil.On("WriteToConsole", mock.Anything, []string{"Title", "Completed"}).
			Return(nil)

		WriteToConsole = mockUtil.WriteToConsole

		app.FetchAndDisplayTodos([]byte("2 4 6 8 10"))
		mockUtil.AssertExpectations(t)
		mockTodoService.AssertExpectations(t)
	})

	t.Run("FetchAndDisplayTodos - Error Writing to CSV", func(t *testing.T) {
		mockUtil := &MockUtil{}
		mockTodoService := &MockTodoService{}

		app := &TodoApp{
			service:    mockTodoService,
			outputFile: "output.csv",
		}

		mockTodoService.On("FetchAndCheckStatus", []int{2, 4, 6, 8, 10}).Return([]todo.Todo{
			{Title: "Todo1", Completed: false},
			{Title: "Todo2", Completed: true},
		})

		mockUtil.On("WriteToCSV", mock.Anything, "output.csv", []string{"Title", "Completed"}).
			Return(errors.New("error writing to CSV"))

		WriteToCSV = mockUtil.WriteToCSV
		app.FetchAndDisplayTodos([]byte("2 4 6 8 10"))
		mockTodoService.AssertExpectations(t)
		mockUtil.AssertExpectations(t)
	})
}

func TestTodoApp_GetFirstNEvenNumbers(t *testing.T) {
	t.Run("GetFirstNEvenNumbers - Successful", func(t *testing.T) {
		result := getFirstNEvenNumbers([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3)
		assert.Equal(t, []int{2, 4, 6}, result)
	})

	t.Run("GetFirstNEvenNumbers - Insufficient Numbers", func(t *testing.T) {
		result := getFirstNEvenNumbers([]int{1, 3, 5}, 3)
		assert.Empty(t, result)
	})
}
