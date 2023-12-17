package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define a sample struct for testing
type SampleStruct struct {
	Name   string
	Age    int
	Active bool
}

func TestWriteToCSV(t *testing.T) {
	// Sample data for testing
	data := []SampleStruct{
		{Name: "John", Age: 30, Active: true},
		{Name: "Alice", Age: 25, Active: false},
	}

	// Positive Test Case
	t.Run("WriteToCSV - Successful Write", func(t *testing.T) {
		// Create a temporary file for testing
		filePath := createTempFilePath()
		defer os.Remove(filePath)

		// Fields to be written to CSV
		fields := []string{"Name", "Age", "Active"}

		err := WriteToCSV(data, filePath, fields)
		assert.NoError(t, err)

		// Read the content of the file to verify
		fileContent, err := readCSVFile(filePath)
		assert.NoError(t, err)

		// Verify CSV content
		expectedContent := "Name,Age,Active\nJohn,30,true\nAlice,25,false\n"
		assert.Equal(t, expectedContent, fileContent)
	})

	// Negative Test Case
	t.Run("WriteToCSV - Invalid File Path", func(t *testing.T) {
		// Specify an invalid file path (e.g., a directory)
		filePath := "/path/to/directory"

		// Fields to be written to CSV
		fields := []string{"Name", "Age", "Active"}

		err := WriteToCSV(data, filePath, fields)
		assert.Error(t, err)
	})
}

func TestWriteToConsole(t *testing.T) {
	// Sample data for testing
	data := []SampleStruct{
		{Name: "John", Age: 30, Active: true},
		{Name: "Alice", Age: 25, Active: false},
	}

	// Positive Test Case
	t.Run("WriteToConsole - Successful Print", func(t *testing.T) {
		// Capture stdout for testing
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Close the write end of the pipe to ensure the reader sees EOF
		defer func() {
			os.Stdout = oldStdout
			w.Close()
		}()

		// Fields to be printed to console
		fields := []string{"Name", "Age", "Active"}

		err := WriteToConsole(data, fields)
		assert.NoError(t, err)

		// Read from the read end of the pipe to get the captured output
		outBytes := make([]byte, 1024)
		n, _ := r.Read(outBytes)

		// Verify console output
		expectedOutput := "Name,Age,Active\nJohn,30,true\nAlice,25,false\n"
		assert.Equal(t, expectedOutput, string(outBytes[:n]))
	})
}

// Helper function to read CSV content from a file
func readCSVFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
