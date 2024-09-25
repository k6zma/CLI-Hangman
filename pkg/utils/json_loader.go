package utils

import (
	"fmt"
	"io"
	"log/slog"
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
		slog.Error("failed to open file", slog.String("filename", filename), slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		slog.Error("failed to read file", slog.String("filename", filename), slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return byteValue, err
}
