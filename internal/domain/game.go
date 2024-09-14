package domain

// GameProperties stores the settings of the game.
type GameProperties struct {
	language    string
	difficulty  string
	maxAttempts int
}

// NewGameProperties initializes a new GameProperties object.
func NewGameProperties(language, difficulty string, maxAttempts int) *GameProperties {
	return &GameProperties{
		language:    language,
		difficulty:  difficulty,
		maxAttempts: maxAttempts,
	}
}

// GetLanguageFromProperties returns the language from the game properties.
func (properties *GameProperties) GetLanguageFromProperties() (string, error) {
	if properties.language == "" {
		return "", NewGetFieldsOfGamePropertiesError("language not set")
	}

	return properties.language, nil
}

// GetDifficultyFromProperties returns the difficulty from the game properties.
func (properties *GameProperties) GetDifficultyFromProperties() (string, error) {
	if properties.difficulty == "" {
		return "", NewGetFieldsOfGamePropertiesError("difficulty not set")
	}

	return properties.difficulty, nil
}

// GetMaxAttemptsFromProperties returns the maximum number of attempts from the game properties.
func (properties *GameProperties) GetMaxAttemptsFromProperties() (int, error) {
	if properties.maxAttempts <= 0 {
		return -1, NewGetFieldsOfGamePropertiesError("max attempts not set")
	}

	return properties.maxAttempts, nil
}

// Game represents the state of the game.
type Game struct {
	word            *Word
	guessedLetters  map[rune]bool
	currentAttempts int
	maxAttempts     int
}

// NewGame initializes a new Game object.
func NewGame(word *Word, maxAttempts int) *Game {
	return &Game{
		word:            word,
		guessedLetters:  make(map[rune]bool),
		currentAttempts: 0,
		maxAttempts:     maxAttempts,
	}
}

// GetWordLetters returns the letters of a word as a rune slice.
func (g *Game) GetWordLetters() ([]rune, error) {
	if g.word == nil {
		return nil, NewGetFieldsOfWordError("word is not set")
	}

	return g.word.GetLetters()
}

// GetWordHint returns a word hint.
func (g *Game) GetWordHint() (string, error) {
	if g.word == nil {
		return "", NewGetFieldsOfWordError("word is not set")
	}

	return g.word.GetHint()
}

// GetMaxAttempts returns the maximum number of attempts.
func (g *Game) GetMaxAttempts() (int, error) {
	if g.maxAttempts == 0 {
		return 0, NewGetFieldsOfGamePropertiesError("max attempts not set")
	}

	return g.maxAttempts, nil
}
