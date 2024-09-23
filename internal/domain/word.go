package domain

// Word represents a word l metadata.
//
// Fields:
// - letters: the runes of the word.
// - language: the language of the word.
// - difficulty: the difficulty of the word.
// - hint: a hint to help the player guess the word.
type Word struct {
	letters    []rune
	language   string
	difficulty string
	hint       string
}

// NewWord initializes a new Word object.
//
// Parameters:
// - wordData: the word data.
// - language: the language of the word.
// - difficulty: the difficulty level.
//
// Function returns:
// - *Word: a pointer to the initialized Word object.
func NewWord(wordData, hint, language, difficulty string) *Word {
	return &Word{
		letters:    []rune(wordData),
		language:   language,
		difficulty: difficulty,
		hint:       hint,
	}
}

// GetLetters returns the letters (runes) of the word.
func (w *Word) GetLetters() ([]rune, error) {
	if len(w.letters) == 0 {
		return nil, NewGetFieldsOfWordError("letters fields are empty")
	}

	return w.letters, nil
}

// GetLanguage returns the language of the word.
func (w *Word) GetLanguage() (*string, error) {
	if w.language == "" {
		return nil, NewGetFieldsOfWordError("language was not set")
	}

	return &w.language, nil
}

// GetDifficulty returns the difficulty of the word.
func (w *Word) GetDifficulty() (*string, error) {
	if w.difficulty == "" {
		return nil, NewGetFieldsOfWordError("difficulty was not set")
	}

	return &w.difficulty, nil
}

// GetHint returns the hint associated with the word.
func (w *Word) GetHint() (*string, error) {
	if w.hint == "" {
		return nil, NewGetFieldsOfWordError("hint was not set")
	}

	return &w.hint, nil
}
