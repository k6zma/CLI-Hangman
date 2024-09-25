package domain

// Сonstants with error texts.
const (
	ErrGetFieldsOfWord           = "error while getting fields of the word"
	ErrGetFieldsOfGameProperties = "error while getting fields of the game properties"
	ErrNewGame                   = "error while initializing a new game"
	ErrGetFieldsOfGame           = "error while getting fields of the game"
)

// GetFieldsOfWordError defines a custom error for issues when retrieving fields from the Word struct.
//
// Fields:
// - message: the error message that describes the issue.
type GetFieldsOfWordError struct {
	message string
}

// NewGetFieldsOfWordError is a constructor that creates and returns a new GetFieldsOfWordError.
//
// Function returns:
// - error: a new instance of GetFieldsOfWordError.
func NewGetFieldsOfWordError() error {
	return &GetFieldsOfWordError{
		message: ErrGetFieldsOfWord,
	}
}

func (err GetFieldsOfWordError) Error() string {
	return err.message
}

// GetFieldsOfGamePropertiesError defines a custom error for issues when retrieving fields from the GameProperties struct.
//
// Fields:
// - message: the error message that describes the issue.
type GetFieldsOfGamePropertiesError struct {
	message string
}

// NewGetFieldsOfGamePropertiesError is a constructor that creates and returns a new GetFieldsOfGamePropertiesError.
//
// Function returns:
// - error: a new instance of GetFieldsOfGamePropertiesError.
func NewGetFieldsOfGamePropertiesError() error {
	return &GetFieldsOfGamePropertiesError{
		message: ErrGetFieldsOfGameProperties,
	}
}

func (err GetFieldsOfGamePropertiesError) Error() string {
	return err.message
}

// NewGameError defines a custom error for issues when initializing a new game.
//
// Fields:
// - message: the error message that describes the issue.
type NewGameError struct {
	message string
}

// MakeNewGameError is a constructor that creates and returns a new NewGameError.
//
// Function returns:
// - error: a new instance of NewGameError.
func MakeNewGameError() error {
	return &NewGameError{
		message: ErrNewGame,
	}
}

func (err NewGameError) Error() string {
	return err.message
}

// GetFieldsOfGameError defines a custom error for issues when retrieving fields from the Game struct.
//
// Fields:
// - message: the error message that describes the issue.
type GetFieldsOfGameError struct {
	message string
}

// NewGetFieldsOfGameError is a constructor that creates and returns a new GetFieldsOfGameError.
//
// Function returns:
// - error: a new instance of GetFieldsOfWordError.
func NewGetFieldsOfGameError() error {
	return &GetFieldsOfGameError{
		message: ErrGetFieldsOfGame,
	}
}

func (err GetFieldsOfGameError) Error() string {
	return err.message
}
