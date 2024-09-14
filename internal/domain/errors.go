package domain

// GetFieldsOfWordError defines a custom error for issues when retrieving fields from the Word struct.
//
// Fields:
// - message: the error message that describes the issue.
type GetFieldsOfWordError struct {
	message string
}

// NewGetFieldsOfWordError is a constructor that creates and returns a new GetFieldsOfWordError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of GetFieldsOfWordError.
func NewGetFieldsOfWordError(message string) error {
	return &GetFieldsOfWordError{
		message: message,
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
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of GetFieldsOfGamePropertiesError.
func NewGetFieldsOfGamePropertiesError(message string) error {
	return &GetFieldsOfGamePropertiesError{
		message: message,
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
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of NewGameError.
func MakeNewGameError(message string) error {
	return &NewGameError{
		message: message,
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
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of GetFieldsOfWordError.
func NewGetFieldsOfGameError(message string) error {
	return &GetFieldsOfGameError{
		message: message,
	}
}

func (err GetFieldsOfGameError) Error() string {
	return err.message
}
