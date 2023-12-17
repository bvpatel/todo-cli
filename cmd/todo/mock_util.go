package todo

import "github.com/stretchr/testify/mock"

type MockUtil struct {
	mock.Mock
}

func (m *MockUtil) ReadFromFile(filePath string) ([]byte, error) {
	args := m.Called(filePath)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockUtil) WriteToCSV(data interface{}, filePath string, fields []string) error {
	args := m.Called(data, filePath, fields)
	return args.Error(0)
}

func (m *MockUtil) WriteToConsole(data interface{}, fields []string) error {
	args := m.Called(data, fields)
	return args.Error(0)
}
