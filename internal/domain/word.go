package domain

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
func NewParsedWords(enWords []WordWithHintJSON, ruWords []WordWithHintJSON) *ParsedWords {
	return &ParsedWords{
		EnWords: enWords,
		RuWords: ruWords,
	}
}

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
// - jsonData: the parsed word data with hints.
// - language: the language of the word.
// - difficulty: the difficulty level.
//
// Function returns:
// - *Word: a pointer to the initialized Word object.
func NewWord(jsonData WordWithHintJSON, language, difficulty string) *Word {
	return &Word{
		letters:    []rune(jsonData.WordData),
		language:   language,
		difficulty: difficulty,
		hint:       jsonData.Hint,
	}
}

// GetLetters returns the letters (runes) of the word.
func (w *Word) GetLetters() ([]rune, error) {
	if len(w.letters) == 0 {
		return nil, NewGetFieldsOfWordError("letters fields is empty")
	}

	return w.letters, nil
}

// GetLanguage returns the language of the word.
func (w *Word) GetLanguage() (string, error) {
	if w.language == "" {
		return "", NewGetFieldsOfWordError("language was not set")
	}

	return w.language, nil
}

// GetDifficulty returns the difficulty of the word.
func (w *Word) GetDifficulty() (string, error) {
	if w.difficulty == "" {
		return "", NewGetFieldsOfWordError("difficulty was not set")
	}

	return w.difficulty, nil
}

// GetHint returns the hint associated with the word.
func (w *Word) GetHint() (string, error) {
	if w.hint == "" {
		return "", NewGetFieldsOfWordError("hint was not set")
	}

	return w.hint, nil
}
