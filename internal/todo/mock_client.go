package todo

import "github.com/stretchr/testify/mock"

type MockTodoClient struct {
	mock.Mock
}

func (m *MockTodoClient) FetchTodo(id int) (Todo, error) {
	args := m.Called(id)
	return args.Get(0).(Todo), args.Error(1)
}
