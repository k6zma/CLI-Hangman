package infrastructure

import (
	"encoding/json"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/utils"
)

var FileReader utils.FileReader = utils.DefaultFileReader{}

// LoadWords loads and parses  words from JSON.
//
// Parameters:
// - filename: the path to the JSON file.
//
// Function returns:
// - *domain.ParsedWords: a pointer to the ParsedWords structure containing the words.
// - error: an error if the file could not be read or the JSON could not be parsed.
func LoadWords(filename string) (*domain.ParsedWords, error) {
	jsonData, err := FileReader.ReadJSON(filename)
	if err != nil {
		return nil, NewWordLoaderError("failed to read JSON")
	}

	var parsedWords domain.ParsedWords
	if err := json.Unmarshal(jsonData, &parsedWords); err != nil {
		return nil, NewWordLoaderError("failed to parse JSON")
	}

	return &parsedWords, nil
}
