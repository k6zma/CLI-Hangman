package domain_test

import (
	"testing"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/stretchr/testify/assert"
)

// --------------------------------
// Stage 1: checking custom errors.
// --------------------------------

// GetFieldsOfWordError error type checking.
func TestNewGetFieldsOfWordError(t *testing.T) {
	err := domain.NewGetFieldsOfWordError("error while getting fields of the word")

	assert.NotNil(t, err)
	assert.Equal(t, "error while getting fields of the word", err.Error())
}

// GetFieldsOfGamePropertiesError error type checking.
func TestNewGetFieldsOfGamePropertiesError(t *testing.T) {
	err := domain.NewGetFieldsOfGamePropertiesError("error while getting fields of the game properties")

	assert.NotNil(t, err)
	assert.Equal(t, "error while getting fields of the game properties", err.Error())
}

// NewGameError error type checking.
func TestMakeNewGameError(t *testing.T) {
	err := domain.MakeNewGameError("new game error")

	assert.NotNil(t, err)
	assert.Equal(t, "new game error", err.Error())
}

// GetFieldsOfGameError error type checking.
func TestNewGetFieldsOfGameError(t *testing.T) {
	err := domain.NewGetFieldsOfGameError("error while getting fields of the game struct")

	assert.NotNil(t, err)
	assert.Equal(t, "error while getting fields of the game struct", err.Error())
}

// --------------------------------
// Stage 2: checking game working.
// --------------------------------

// Creating game with english word.
func TestNewGameEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "easy")
	game := domain.NewGame(word, 5)

	assert.NotNil(t, game)

	maxAttempts, err := game.GetMaxAttempts()

	assert.NoError(t, err)
	assert.Equal(t, 5, *maxAttempts)
}

// Creating game with russian word.
func TestNewGameRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "easy",
	)

	game := domain.NewGame(word, 5)

	assert.NotNil(t, game)

	maxAttempts, err := game.GetMaxAttempts()

	assert.NoError(t, err)
	assert.Equal(t, 5, *maxAttempts)
}

// -------------------------------
// Stage 2: checking game getters.
// -------------------------------

// Validation check for the correctness of the entered letter in english.
func TestGuessLetterEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "easy")
	game := domain.NewGame(word, 5)

	assert.True(t, game.GuessLetter('l'))
	assert.False(t, game.GuessLetter('a'))
}

// Validation check for the correctness of the entered letter in russian.
func TestGuessLetterRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "easy",
	)

	game := domain.NewGame(word, 5)

	assert.True(t, game.GuessLetter('л'))
	assert.False(t, game.GuessLetter('б'))
}

// Checking the correctness of the storage of letters in english.
func TestGetWordLettersEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "easy")
	game := domain.NewGame(word, 5)

	letters, err := game.GetWordLetters()

	assert.NoError(t, err)
	assert.Equal(t, []rune{'l', 'i', 'm', 'e'}, letters)
}

// Checking the correctness of the storage of letters in russian.
func TestGetWordLettersRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "easy",
	)

	game := domain.NewGame(word, 5)

	letters, err := game.GetWordLetters()

	assert.NoError(t, err)
	assert.Equal(t, []rune{'л', 'а', 'й', 'м'}, letters)
}

// -------------------------------------
// Stage 3: checking game logic methods.
// -------------------------------------

// Verifying that game end validation is correct in english.
func TestIsGameOverEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "easy")
	game := domain.NewGame(word, 2)

	game.GuessLetter('a')
	game.GuessLetter('b')

	assert.True(t, game.IsGameOver())
}

// Verifying that game end validation is correct in russian.
func TestIsGameOverRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "easy",
	)

	game := domain.NewGame(word, 2)

	game.GuessLetter('б')
	game.GuessLetter('в')

	assert.True(t, game.IsGameOver())
}

// Verifying that the game win is validated correctly in english.
func TestIsGameWonEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "easy")
	game := domain.NewGame(word, 5)

	game.GuessLetter('l')
	game.GuessLetter('i')
	game.GuessLetter('m')
	game.GuessLetter('e')

	assert.True(t, game.IsGameWon())
}

// Verifying that the game win is validated correctly in russian.
func TestIsGameWonRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "easy",
	)

	game := domain.NewGame(word, 5)

	game.GuessLetter('л')
	game.GuessLetter('а')
	game.GuessLetter('й')
	game.GuessLetter('м')

	assert.True(t, game.IsGameWon())
}

// Checking whether the word hint is correct in english.
func TestGetWordHintEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "easy")
	game := domain.NewGame(word, 5)

	hint, err := game.GetWordHint()

	assert.NoError(t, err)
	assert.Equal(t, "green citrus fruit used in cocktails and cooking", *hint)
}

// Checking whether the word hint is correct in russian.
func TestGetWordHintRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "easy",
	)

	game := domain.NewGame(word, 5)

	hint, err := game.GetWordHint()

	assert.NoError(t, err)
	assert.Equal(t, "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии", *hint)
}

// Checking for correct storage of already guessed letters in english.
func TestGetGuessedLettersEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "easy")
	game := domain.NewGame(word, 5)

	game.GuessLetter('l')

	guessedLetters, err := game.GetGuessedLetters()

	assert.NoError(t, err)

	assert.True(t, guessedLetters['l'])
	assert.False(t, guessedLetters['m'])
}

// Checking for correct storage of already guessed letters in russian.
func TestGetGuessedLettersRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "easy",
	)

	game := domain.NewGame(word, 5)

	game.GuessLetter('л')

	guessedLetters, err := game.GetGuessedLetters()

	assert.NoError(t, err)

	assert.True(t, guessedLetters['л'])
	assert.False(t, guessedLetters['м'])
}

// ----------------------------------------------------
// Stage 4: checking constructors for WordWithHintJSON.
// ----------------------------------------------------

// Checks that the constructor worked and the object is not empty in english.
func TestNewWordWithHintJSONEn(t *testing.T) {
	wordWithHint := domain.NewWordWithHintJSON("lime", "green citrus fruit used in cocktails and cooking")

	assert.NotNil(t, wordWithHint)
	assert.Equal(t, "lime", wordWithHint.WordData)
	assert.Equal(t, "green citrus fruit used in cocktails and cooking", wordWithHint.Hint)
}

// Checks that the constructor worked and the object is not empty in russian.
func TestNewWordWithHintJSONRu(t *testing.T) {
	wordWithHint := domain.NewWordWithHintJSON("лайм", "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии")

	assert.NotNil(t, wordWithHint)
	assert.Equal(t, "лайм", wordWithHint.WordData)
	assert.Equal(t, "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии", wordWithHint.Hint)
}

// -----------------------------------------------
// Stage 5: checking constructors for ParsedWords.
// -----------------------------------------------

// Checking the functionality of the NewParsedWords constructor.
func TestNewParsedWords(t *testing.T) {
	enWords := []domain.WordWithHintJSON{
		{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"},
	}

	ruWords := []domain.WordWithHintJSON{
		{WordData: "лайм", Hint: "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
	}

	parsedWords := domain.NewParsedWords(enWords, ruWords)

	assert.NotNil(t, parsedWords)
	assert.Equal(t, enWords, parsedWords.EnWords)
	assert.Equal(t, ruWords, parsedWords.RuWords)
}

// ----------------------------------------
// Stage 6: checking constructors for Word.
// ----------------------------------------

// Multilayer NewWord constructor validation in english.
func TestNewWordEn(t *testing.T) {
	wordWithHint := domain.NewWordWithHintJSON("lime", "green citrus fruit used in cocktails and cooking")
	word := domain.NewWord(*wordWithHint, "en", "easy")

	assert.NotNil(t, word)

	letters, err := word.GetLetters()

	assert.NoError(t, err)
	assert.Equal(t, []rune{'l', 'i', 'm', 'e'}, letters)

	language, err := word.GetLanguage()

	assert.NoError(t, err)
	assert.Equal(t, "en", *language)

	difficulty, err := word.GetDifficulty()

	assert.NoError(t, err)
	assert.Equal(t, "easy", *difficulty)

	hint, err := word.GetHint()

	assert.NoError(t, err)
	assert.Equal(t, "green citrus fruit used in cocktails and cooking", *hint)
}

// Multilayer NewWord constructor validation in russian.
func TestNewWordRu(t *testing.T) {
	wordWithHint := domain.NewWordWithHintJSON("лайм", "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии")
	word := domain.NewWord(*wordWithHint, "ru", "easy")

	assert.NotNil(t, word)

	letters, err := word.GetLetters()

	assert.NoError(t, err)
	assert.Equal(t, []rune{'л', 'а', 'й', 'м'}, letters)

	language, err := word.GetLanguage()

	assert.NoError(t, err)
	assert.Equal(t, "ru", *language)

	difficulty, err := word.GetDifficulty()

	assert.NoError(t, err)
	assert.Equal(t, "easy", *difficulty)

	hint, err := word.GetHint()

	assert.NoError(t, err)
	assert.Equal(t, "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии", *hint)
}

// -------------------------------------
// Stage 7: checking getters validation.
// -------------------------------------

// Checking the operation of the validator in the getter if the word is not specified or empty in english.
func TestGetLettersEmptyWordEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "", Hint: "empty word"}, "en", "easy")

	letters, err := word.GetLetters()

	assert.Nil(t, letters)
	assert.Error(t, err)
	assert.Equal(t, "letters fields are empty", err.Error())
}

// Checking the operation of the validator in the getter if the word is not specified or empty in russian.
func TestGetLettersEmptyWordRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "", Hint: "пустое слово"}, "ru", "easy")

	letters, err := word.GetLetters()

	assert.Nil(t, letters)
	assert.Error(t, err)
	assert.Equal(t, "letters fields are empty", err.Error())
}

// Checking the operation of the validator in the getter if the language is not specified or is empty in english.
func TestGetLanguageEmptyEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "", "easy")

	language, err := word.GetLanguage()

	assert.Empty(t, language)
	assert.Error(t, err)
	assert.Equal(t, "language was not set", err.Error())
}

// Checking the operation of the validator in the getter if the language is not specified or is empty in russian.
func TestGetLanguageEmptyRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"", "easy",
	)

	language, err := word.GetLanguage()

	assert.Empty(t, language)
	assert.Error(t, err)
	assert.Equal(t, "language was not set", err.Error())
}

// Checking the operation of the validator in the getter if the complexity is not specified or is empty in english.
func TestGetDifficultyEmptyEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: "green citrus fruit used in cocktails and cooking"}, "en", "")

	difficulty, err := word.GetDifficulty()

	assert.Empty(t, difficulty)
	assert.Error(t, err)
	assert.Equal(t, "difficulty was not set", err.Error())
}

// Checking the operation of the validator in the getter if the complexity is not specified or is empty in russian.
func TestGetDifficultyEmptyRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{
		WordData: "лайм",
		Hint:     "зеленый цитрусовый фрукт используемый в коктейлях и кулинарии"},
		"ru", "",
	)

	difficulty, err := word.GetDifficulty()

	assert.Empty(t, difficulty)
	assert.Error(t, err)
	assert.Equal(t, "difficulty was not set", err.Error())
}

// Checking the operation of the validator in the getter if the hint is not specified or is empty in english.
func TestGetHintEmptyEn(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "lime", Hint: ""}, "en", "easy")

	hint, err := word.GetHint()

	assert.Empty(t, hint)
	assert.Error(t, err)
	assert.Equal(t, "hint was not set", err.Error())
}

// Checking the operation of the validator in the getter if the hint is not specified or is empty in russian.
func TestGetHintEmptyRu(t *testing.T) {
	word := domain.NewWord(domain.WordWithHintJSON{WordData: "лайм", Hint: ""}, "ru", "easy")

	hint, err := word.GetHint()

	assert.Empty(t, hint)
	assert.Error(t, err)
	assert.Equal(t, "hint was not set", err.Error())
}
