package application

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"

	"crypto/rand"
	"math/big"
)

// SelectWordByDifficulty filters words based on difficulty and returns a random word.
//
// Parameters:
// - words: a slice of WordWithHintJSON containing available words.
// - difficulty: the selected difficulty level.
//
// Function returns:
// - *domain.WordWithHintJSON: a pointer to the selected word.
// - error: an error if no word matches the difficulty or a bad difficulty is entered.
func SelectWordByDifficulty(words []domain.WordWithHintJSON, difficulty string) (*domain.WordWithHintJSON, error) {
	var filteredWords []domain.WordWithHintJSON

	for _, word := range words {
		switch difficulty {
		case "easy":
			if len(word.WordData) <= 4 {
				filteredWords = append(filteredWords, word)
			}
		case "medium":
			if len(word.WordData) > 4 && len(word.WordData) <= 7 {
				filteredWords = append(filteredWords, word)
			}
		case "hard":
			if len(word.WordData) > 7 {
				filteredWords = append(filteredWords, word)
			}
		default:
			return nil, NewWordSelectorError("invalid difficulty level")
		}
	}

	if len(filteredWords) == 0 {
		return nil, NewWordSelectorError("no words found for the selected difficulty")
	}

	maxIndex := big.NewInt(int64(len(filteredWords)))
	randomIndex, err := rand.Int(rand.Reader, maxIndex)

	if err != nil {
		return nil, NewWordSelectorError("error generating random index")
	}

	return &filteredWords[randomIndex.Int64()], nil
}
