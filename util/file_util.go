package util

import (
	"io/ioutil"
	"os"
)

// ReadFromFile reads data from the specified file
func ReadFromFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

// WriteToFile writes data to the specified file
func WriteToFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
