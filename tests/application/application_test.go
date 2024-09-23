package application_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

// --------------------------------
// Stage 1: checking custom errors.
// --------------------------------

// WordSelectorError error type checking.
func TestWordSelectorError(t *testing.T) {
	err := application.NewWordSelectorError("failed to select the word")
	assert.NotNil(t, err)
	assert.Equal(t, "failed to select the word", err.Error())
}

// WordsLoadingError error type checking.
func TestWordsLoadingError(t *testing.T) {
	err := application.NewWordsLoadingError("failed to load words")
	assert.NotNil(t, err)
	assert.Equal(t, "failed to load words", err.Error())
}

// GamePropertiesCollectingError error type checking.
func TestGamePropertiesCollectingError(t *testing.T) {
	err := application.NewGamePropertiesCollectingError("failed to collect game properties")
	assert.NotNil(t, err)
	assert.Equal(t, "failed to collect game properties", err.Error())
}

// --------------------------------
// Stage 2: checking word selection by difficulty and language.
// --------------------------------

// Сhecking the game service constructor.
func TestNewGameService(t *testing.T) {
	gameService := application.NewGameService()
	assert.NotNil(t, gameService)
}

// Checking word selection using user properties.
//
// Properties:
// - Language: en.
// - Difficulty: easy.
// - MaxAttempts: 5.
func TestSelectWordByPropertiesEnEasy(t *testing.T) {
	gameService := application.NewGameService()

	words := infrastructure.ParsedWords{
		EnWords: []infrastructure.WordWithHintJSON{
			{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"},
			{WordData: "lemon", Hint: "yellow citrus fruit used in cocktails and cooking"},
			{WordData: "tangerine", Hint: "orange small citrus fruit"},
		},
	}

	gameProperties := domain.NewGameProperties("en", "easy", 5)

	selectedWord, err := gameService.SelWordByPr(&words, gameProperties)

	assert.NoError(t, err)
	assert.NotNil(t, selectedWord)
	assert.Equal(t, "lime", selectedWord.WordData)
}

// Checking word selection using user properties.
//
// Properties:
// - Language: en.
// - Difficulty: medium.
// - MaxAttempts: 5.
func TestSelectWordByPropertiesEnMedium(t *testing.T) {
	gameService := application.NewGameService()

	words := infrastructure.ParsedWords{
		EnWords: []infrastructure.WordWithHintJSON{
			{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"},
			{WordData: "lemon", Hint: "yellow citrus fruit used in cocktails and cooking"},
			{WordData: "tangerine", Hint: "orange small citrus fruit"},
		},
	}

	gameProperties := domain.NewGameProperties("en", "medium", 5)

	selectedWord, err := gameService.SelWordByPr(&words, gameProperties)

	assert.NoError(t, err)
	assert.NotNil(t, selectedWord)
	assert.Equal(t, "lemon", selectedWord.WordData)
}

// Checking word selection using user properties.
//
// Properties:
// - Language: en.
// - Difficulty: hard.
// - MaxAttempts: 5.
func TestSelectWordByPropertiesEnHard(t *testing.T) {
	gameService := application.NewGameService()

	words := infrastructure.ParsedWords{
		EnWords: []infrastructure.WordWithHintJSON{
			{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"},
			{WordData: "lemon", Hint: "yellow citrus fruit used in cocktails and cooking"},
			{WordData: "tangerine", Hint: "orange small citrus fruit"},
		},
	}

	gameProperties := domain.NewGameProperties("en", "hard", 5)

	selectedWord, err := gameService.SelWordByPr(&words, gameProperties)

	assert.NoError(t, err)
	assert.NotNil(t, selectedWord)
	assert.Equal(t, "tangerine", selectedWord.WordData)
}

// Checking word selection using user properties.
//
// Properties:
// - Language: ru.
// - Difficulty: easy.
// - MaxAttempts: 5.
func TestSelectWordByPropertiesRuEasy(t *testing.T) {
	gameService := application.NewGameService()

	words := infrastructure.ParsedWords{
		RuWords: []infrastructure.WordWithHintJSON{
			{WordData: "лайм", Hint: "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
			{WordData: "лимон", Hint: "желтый цитрусовый фрукт используемый в коктейлях и кулинарии"},
			{WordData: "мандарин", Hint: "оранжевый маленький цитрусовый фрукт"},
		},
	}

	gameProperties := domain.NewGameProperties("ru", "easy", 5)

	selectedWord, err := gameService.SelWordByPr(&words, gameProperties)

	assert.NoError(t, err)
	assert.NotNil(t, selectedWord)
	assert.Equal(t, "лайм", selectedWord.WordData)
}

// Checking word selection using user properties.
//
// Properties:
// - Language: ru.
// - Difficulty: medium.
// - MaxAttempts: 5.
func TestSelectWordByPropertiesRuMedium(t *testing.T) {
	gameService := application.NewGameService()

	words := infrastructure.ParsedWords{
		RuWords: []infrastructure.WordWithHintJSON{
			{WordData: "лайм", Hint: "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
			{WordData: "лимон", Hint: "желтый цитрусовый фрукт используемый в коктейлях и кулинарии"},
			{WordData: "мандарин", Hint: "оранжевый маленький цитрусовый фрукт"},
		},
	}

	gameProperties := domain.NewGameProperties("ru", "medium", 5)

	selectedWord, err := gameService.SelWordByPr(&words, gameProperties)

	assert.NoError(t, err)
	assert.NotNil(t, selectedWord)
	assert.Equal(t, "лимон", selectedWord.WordData)
}

// Checking word selection using user properties.
//
// Properties:
// - Language: ru.
// - Difficulty: hard.
// - MaxAttempts: 5.
func TestSelectWordByPropertiesRuHard(t *testing.T) {
	gameService := application.NewGameService()

	words := infrastructure.ParsedWords{
		RuWords: []infrastructure.WordWithHintJSON{
			{WordData: "лайм", Hint: "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
			{WordData: "лимон", Hint: "желтый цитрусовый фрукт используемый в коктейлях и кулинарии"},
			{WordData: "мандарин", Hint: "оранжевый маленький цитрусовый фрукт"},
		},
	}

	gameProperties := domain.NewGameProperties("ru", "hard", 5)

	selectedWord, err := gameService.SelWordByPr(&words, gameProperties)

	assert.NoError(t, err)
	assert.NotNil(t, selectedWord)
	assert.Equal(t, "мандарин", selectedWord.WordData)
}

// Checking word selection using wrong difficulty.
//
// Properties:
// - Language: en.
// - Difficulty: impossible.
func TestInvalidDifficultyEn(t *testing.T) {
	words := []infrastructure.WordWithHintJSON{
		{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"},
	}

	selectedWord, err := application.SelectWordByDifficulty(words, "impossible")
	assert.Error(t, err)
	assert.Nil(t, selectedWord)
	assert.EqualError(t, err, "invalid difficulty level")
}

// Checking word selection using wrong difficulty.
//
// Properties:
// - Language: en.
// - Difficulty: невозможно.
func TestInvalidDifficultyRu(t *testing.T) {
	words := []infrastructure.WordWithHintJSON{
		{WordData: "лайм", Hint: "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
	}

	selectedWord, err := application.SelectWordByDifficulty(words, "невозомжно")
	assert.Error(t, err)
	assert.Nil(t, selectedWord)
	assert.EqualError(t, err, "invalid difficulty level")
}

// Checking word selection where there are no words for the required complexity.
//
// Properties:
// - Language: en.
// - Difficulty: hard.
func TestNoWordsForDifficultyEn(t *testing.T) {
	words := []infrastructure.WordWithHintJSON{
		{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"},
	}

	selectedWord, err := application.SelectWordByDifficulty(words, "hard")
	assert.Error(t, err)
	assert.Nil(t, selectedWord)
	assert.EqualError(t, err, "no words found for the selected difficulty")
}

// Checking word selection where there are no words for the required complexity.
//
// Properties:
// - Language: ru.
// - Difficulty: hard.
func TestNoWordsForDifficultyRu(t *testing.T) {
	words := []infrastructure.WordWithHintJSON{
		{WordData: "лайм", Hint: "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
	}

	selectedWord, err := application.SelectWordByDifficulty(words, "hard")
	assert.Error(t, err)
	assert.Nil(t, selectedWord)
	assert.EqualError(t, err, "no words found for the selected difficulty")
}

// --------------------------------
// Stage 3: checking game session.
// --------------------------------

// Checking that the game session is actually starting and is not nil.
func TestStartGameSession(t *testing.T) {
	gameService := application.NewGameService()

	selectedWord := infrastructure.WordWithHintJSON{
		WordData: "lime",
		Hint:     "green citrus fruit used in cocktails and cooking",
	}

	gameProperties := domain.NewGameProperties("en", "easy", 5)

	game, err := gameService.StartGameSession(&selectedWord, gameProperties)
	assert.NoError(t, err)
	assert.NotNil(t, game)
}
