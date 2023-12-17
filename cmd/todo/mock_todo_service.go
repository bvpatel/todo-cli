package todo

import (
	"todo-cli/internal/todo"

	"github.com/stretchr/testify/mock"
)

type MockTodoService struct {
	mock.Mock
}

func (m *MockTodoService) FetchAndCheckStatus(numbers []int) []todo.Todo {
	args := m.Called(numbers)
	return args.Get(0).([]todo.Todo)
}
