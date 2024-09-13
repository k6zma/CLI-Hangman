package infrastructure

import (
	"encoding/json"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/utils"
)

func LoadWords(filename string) (*domain.ParsedWords, error) {
	jsonData, err := utils.ReadJSON(filename)
	if err != nil {
		return nil, domain.NewWordLoaderError("failed to read JSON")
	}

	var parsedWords domain.ParsedWords
	if err := json.Unmarshal(jsonData, &parsedWords); err != nil {
		return nil, domain.NewWordLoaderError("failed to parse JSON")
	}

	return &parsedWords, nil
}
