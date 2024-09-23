package infrastructure

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/es-debug/backend-academy-2024-go-template/pkg/utils"
)

// WordWithHintJSON represents a structure for a word and its hint parsed from JSON.
//
// Fields:
// - WordData: The word to guess.
// - Hint: A hint to help the player guess the word.
type WordWithHintJSON struct {
	WordData string `json:"word"`
	Hint     string `json:"hint"`
}

// NewWordWithHintJSON initializes a new WordWithHintJSON object.
//
// Parameters:
// - word: the word to guess.
// - hint: a hint to help the player guess the word.
//
// Function returns:
// - *WordWithHintJSON: a pointer to the initialized WordWithHintJSON object.
func NewWordWithHintJSON(word, hint string) *WordWithHintJSON {
	return &WordWithHintJSON{
		WordData: word,
		Hint:     hint,
	}
}

// ParsedWords represents the structure which has words in different languages parsed from JSON.
//
// Fields:
// - EnWords: English words with hints.
// - RuWords: Russian words with hints.
type ParsedWords struct {
	EnWords []WordWithHintJSON `json:"en-words"`
	RuWords []WordWithHintJSON `json:"ru-words"`
}

// NewParsedWords initializes a new ParsedWords object.
//
// Parameters:
// - enWords: a slice of WordWithHintJSON representing English words with hints.
// - ruWords: a slice of WordWithHintJSON representing Russian words with hints.
//
// Function returns:
// - *ParsedWords: a pointer to the initialized ParsedWords object.
func NewParsedWords(enWords, ruWords []WordWithHintJSON) *ParsedWords {
	return &ParsedWords{
		EnWords: enWords,
		RuWords: ruWords,
	}
}

// LoadWords loads and parses  words from JSON.
//
// Parameters:
// - filename: the path to the JSON file.
//
// Function returns:
// - *domain.ParsedWords: a pointer to the ParsedWords structure containing the words.
// - error: an error if the file could not be read or the JSON could not be parsed.
func LoadWords(filename string) (*ParsedWords, error) {
	jsonData, err := utils.ReadJSON(filename)
	if err != nil {
		slog.Error("failed to read json file", slog.String("filename", filename), slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to read json: %w", err)
	}

	var parsedWords ParsedWords
	if err := json.Unmarshal(jsonData, &parsedWords); err != nil {
		slog.Error("failed to unmarshal json", slog.String("filename", filename), slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to parse(unmarshall) json: %w", err)
	}

	return &parsedWords, nil
}
