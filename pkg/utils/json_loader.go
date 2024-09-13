package utils

import (
	"io"
	"os"
)

// ReadJSON opens a file and reads its data.
//
// Parameters:
// - filename: a string representing the path to the JSON file.
//
// Function returns:
// - []byte: a byte slice containing the file content.
// - error: an error if there was an issue opening or reading the file.
func ReadJSON(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return byteValue, err
}
