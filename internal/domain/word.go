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

// ParsedWords represents the structure which has words in different languages parsed from JSON.
//
// Fields:
// - EnWords: English words with hints.
// - RuWords: Russian words with hints.
type ParsedWords struct {
	EnWords []WordWithHintJSON `json:"en-words"`
	RuWords []WordWithHintJSON `json:"ru-words"`
}
