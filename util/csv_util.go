package util

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// WriteToCSV writes a list of structs to a CSV file with specified fields
func WriteToCSV(data interface{}, filePath string, fields []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	if err := writer.Write(fields); err != nil {
		return err
	}

	// Write data to CSV
	val := reflect.ValueOf(data)
	for i := 0; i < val.Len(); i++ {
		var row []string
		for _, field := range fields {
			// Get the value of the field from the struct
			fieldValue := val.Index(i).FieldByName(field).Interface()
			// Convert the field value to string and append to the row
			row = append(row, convertToString(fieldValue))
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

// convertToString converts a value to its string representation
func convertToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}

// writeToConsole prints CSV content to console
func WriteToConsole(data interface{}, fields []string) error {
	val := reflect.ValueOf(data)

	// Print CSV header to console
	fmt.Println(strings.Join(fields, ","))

	// Print data to console
	for i := 0; i < val.Len(); i++ {
		var row []string
		for _, field := range fields {
			// Get the value of the field from the struct
			fieldValue := val.Index(i).FieldByName(field).Interface()
			// Convert the field value to string and append to the row
			row = append(row, convertToString(fieldValue))
		}
		fmt.Println(strings.Join(row, ","))
	}

	return nil
}
