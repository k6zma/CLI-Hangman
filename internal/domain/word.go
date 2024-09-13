package domain

// ParsedWords represents the structure which has words in different languages parsed from JSON
//
// Fields:
// - EnWords: English words
// - RuWords: Russian words
type ParsedWords struct {
	EnWords []string `json:"en-words"`
	RuWords []string `json:"ru-words"`
}
