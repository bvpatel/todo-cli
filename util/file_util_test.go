package util

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromFile(t *testing.T) {
	// Positive Test Case
	t.Run("ReadFromFile - Existing File", func(t *testing.T) {
		// Create a temporary file for testing
		content := []byte("Hello, World!")
		filePath := createTempFile(content)
		defer os.Remove(filePath)

		readContent, err := ReadFromFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, content, readContent)
	})

	// Negative Test Case
	t.Run("ReadFromFile - Nonexistent File", func(t *testing.T) {
		// Specify a non-existent file path
		filePath := "/path/to/nonexistent/file.txt"

		_, err := ReadFromFile(filePath)
		assert.Error(t, err)
	})
}

func TestWriteToFile(t *testing.T) {
	// Positive Test Case
	t.Run("WriteToFile - Successful Write", func(t *testing.T) {
		// Create a temporary file path
		filePath := createTempFilePath()
		defer os.Remove(filePath)

		// Data to be written
		data := []byte("Hello, World!")

		err := WriteToFile(filePath, data)
		assert.NoError(t, err)

		// Read the content of the file to verify
		readContent, err := ioutil.ReadFile(filePath)
		assert.NoError(t, err)
		assert.Equal(t, data, readContent)
	})

	// Negative Test Case
	t.Run("WriteToFile - Invalid File Path", func(t *testing.T) {
		// Specify an invalid file path (e.g., a directory)
		filePath := "/path/to/directory"

		// Data to be written
		data := []byte("Hello, World!")

		err := WriteToFile(filePath, data)
		assert.Error(t, err)
	})
}

// Helper function to create a temporary file with content and return its path
func createTempFile(content []byte) string {
	file, err := ioutil.TempFile("", "testfile-")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = ioutil.WriteFile(file.Name(), content, 0644)
	if err != nil {
		panic(err)
	}

	return file.Name()
}

// Helper function to create a temporary file path and return it
func createTempFilePath() string {
	file, err := ioutil.TempFile("", "testfile-")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return file.Name()
}
