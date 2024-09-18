package infrastructure_test

import (
	"os"
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

// --------------------------------
// Stage 1: checking custom errors.
// --------------------------------

// WordLoaderError error type checking.
func TestNewWordLoaderError(t *testing.T) {
	err := infrastructure.NewWordLoaderError("failed to load words")
	assert.EqualError(t, err, "failed to load words")
}

// InputLanguageError error type checking.
func TestNewInputLanguageError(t *testing.T) {
	err := infrastructure.NewInputLanguageError("invalid language")
	assert.EqualError(t, err, "invalid language")
}

// InputDifficultyError error type checking.
func TestNewInputDifficultyError(t *testing.T) {
	err := infrastructure.NewInputDifficultyError("invalid difficulty")
	assert.EqualError(t, err, "invalid difficulty")
}

// InputMaxAttemptsError error type checking.
func TestNewInputMaxAttemptsError(t *testing.T) {
	err := infrastructure.NewInputMaxAttemptsError("invalid max attempts")
	assert.EqualError(t, err, "invalid max attempts")
}

// InputLetterError error type checking.
func TestNewInputLetterError(t *testing.T) {
	err := infrastructure.NewInputLetterError("invalid letter")
	assert.EqualError(t, err, "invalid letter")
}

// --------------------------------
// Stage 2: checking word loading.
// --------------------------------

// TestLoadWordsSuccess checks the success of loading words from a tmp file.
func TestLoadWordsSuccess(t *testing.T) {
	content := `{
		"en-words": [{"word": "lime", "hint": "green citrus fruit used in cocktails and cooking"}],
		"ru-words": [{"word": "лайм", "hint": "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"}]
	}`

	tmpFile, err := os.CreateTemp("", "words.json")

	assert.NoError(t, err)

	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)

	_, err = tmpFile.Seek(0, 0)
	assert.NoError(t, err)

	wordsData, err := infrastructure.LoadWords(tmpFile.Name())

	assert.NoError(t, err)
	assert.NotNil(t, wordsData)

	assert.Equal(t, "lime", wordsData.EnWords[0].WordData)
	assert.Equal(t, "лайм", wordsData.RuWords[0].WordData)
}

// TestLoadWordsFail checks the fail of loading words from a tmp file.
func TestLoadWordsFail(t *testing.T) {
	wordsData, err := infrastructure.LoadWords("invalid.json")

	assert.Error(t, err)
	assert.Nil(t, wordsData)
	assert.Contains(t, err.Error(), "failed to read JSON")
}
