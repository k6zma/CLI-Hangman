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
func (properties *GameProperties) GetLanguageFromProperties() (*string, error) {
	if properties.language == "" {
		return nil, NewGetFieldsOfGamePropertiesError("language not set")
	}

	return &properties.language, nil
}

// GetDifficultyFromProperties returns the difficulty from the game properties.
func (properties *GameProperties) GetDifficultyFromProperties() (*string, error) {
	if properties.difficulty == "" {
		return nil, NewGetFieldsOfGamePropertiesError("difficulty not set")
	}

	return &properties.difficulty, nil
}

// GetMaxAttemptsFromProperties returns the maximum number of attempts from the game properties.
func (properties *GameProperties) GetMaxAttemptsFromProperties() (*int, error) {
	if properties.maxAttempts <= 0 {
		return nil, NewGetFieldsOfGamePropertiesError("max attempts not set")
	}

	return &properties.maxAttempts, nil
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
func (g *Game) GetWordHint() (*string, error) {
	if g.word == nil {
		return nil, NewGetFieldsOfWordError("word is not set")
	}

	return g.word.GetHint()
}

// GetCurrentAttempts returns the current number of incorrect attempts.
func (g *Game) GetCurrentAttempts() (*int, error) {
	if g.currentAttempts < 0 {
		return nil, NewGetFieldsOfGameError("current attempts not set or invalid")
	}

	return &g.currentAttempts, nil
}

// GetMaxAttempts returns the maximum number of attempts.
func (g *Game) GetMaxAttempts() (*int, error) {
	if g.maxAttempts == 0 {
		return nil, NewGetFieldsOfGameError("max attempts not set")
	}

	return &g.maxAttempts, nil
}

// GetGuessedLetters returns the map of guessed letters or an error if the map is not initialized.
func (g *Game) GetGuessedLetters() (map[rune]bool, error) {
	if g.guessedLetters == nil {
		return nil, NewGetFieldsOfGameError("guessed letters map is not initialized")
	}

	return g.guessedLetters, nil
}

// GetLanguage returns the language of the word in the game.
func (g *Game) GetLanguage() (*string, error) {
	if g.word == nil {
		return nil, NewGetFieldsOfWordError("word is not set")
	}

	language, err := g.word.GetLanguage()
	if err != nil {
		return nil, err
	}

	return language, nil
}

// IsGameOver checks if the game is over based on the number of attempts.
//
// Function returns:
// - bool: true if the game is over, false otherwise.
func (g *Game) IsGameOver() bool {
	return g.currentAttempts >= g.maxAttempts
}

// IsGameWon checks if the game is won by verifying all letters have been guessed.
//
// Function returns:
// - bool: true if the game is won, false otherwise.
func (g *Game) IsGameWon() bool {
	for _, letter := range g.word.letters {
		if !g.guessedLetters[letter] {
			return false
		}
	}

	return true
}

// GuessLetter processes a guessed letter and updates the game state.
//
// Parameters:
// - letter: the rune representing the guessed letter.
//
// Function returns:
// - bool: true if the guessed letter is correct, false otherwise.
func (g *Game) GuessLetter(letter rune) bool {
	g.guessedLetters[letter] = true
	isCorrect := false

	for _, wordLetter := range g.word.letters {
		if wordLetter == letter {
			isCorrect = true
			break
		}
	}

	if !isCorrect {
		g.currentAttempts++
	}

	return isCorrect
}
