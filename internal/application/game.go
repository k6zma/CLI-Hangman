package application

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
)

type GameService struct{}

// StartGame initializes a new game.
//
// Parameters:
// - wordsFilename: a string representing the path to the JSON file containing the word list.
//
// Function returns:
// - *domain.Game: a pointer to the initialized Game struct.
// - error: an error if there was an issue loading words, selecting game properties, or selecting a word.
func (s *GameService) StartGame(wordsFilename string) (*domain.Game, error) {
	wordsData, err := infrastructure.LoadWords(wordsFilename)
	if err != nil {
		return nil, err
	}

	infrastructure.HangmanIntro()

	gameProperties := infrastructure.RequestGameProperties()

	language, err := gameProperties.GetLanguageFromProperties()
	if err != nil {
		return nil, err
	}

	difficulty, err := gameProperties.GetDifficultyFromProperties()
	if err != nil {
		return nil, err
	}

	maxAttempts, err := gameProperties.GetMaxAttemptsFromProperties()
	if err != nil {
		return nil, err
	}

	var selectedWords []domain.WordWithHintJSON

	switch language {
	case "en":
		selectedWords = wordsData.EnWords
	case "ru":
		selectedWords = wordsData.RuWords
	}

	selectedWord, err := SelectWordByDifficulty(selectedWords, difficulty)
	if err != nil {
		return nil, err
	}

	word := domain.NewWord(*selectedWord, language, difficulty)
	game := domain.NewGame(word, maxAttempts)

	return game, nil
}
