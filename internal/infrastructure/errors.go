package infrastructure

// WordLoaderError defines a custom error of word loading.
//
// Fields:
// - message: the error message that describes the issue.
type WordLoaderError struct {
	message string
}

// NewWordLoaderError is constructor which creates and returns a new WordLoaderError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of WordLoaderError.
func NewWordLoaderError(message string) error {
	return &WordLoaderError{
		message: message,
	}
}

func (err WordLoaderError) Error() string {
	return err.message
}

// InputLanguageError defines a custom error for invalid language input.
//
// Fields:
// - message: the error message that describes the issue.
type InputLanguageError struct {
	message string
}

// NewInputLanguageError is a constructor that creates and returns a new InputLanguageError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of InputLanguageError.
func NewInputLanguageError(message string) error {
	return &InputLanguageError{
		message: message,
	}
}

func (err InputLanguageError) Error() string {
	return err.message
}

// InputDifficultyError defines a custom error for invalid difficulty input.
//
// Fields:
// - message: the error message that describes the issue.
type InputDifficultyError struct {
	message string
}

// NewInputDifficultyError is a constructor that creates and returns a new InputDifficultyError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of InputDifficultyError.
func NewInputDifficultyError(message string) error {
	return &InputDifficultyError{
		message: message,
	}
}

func (err InputDifficultyError) Error() string {
	return err.message
}

// InputMaxAttemptsError defines a custom error for invalid max attempts input.
//
// Fields:
// - message: the error message that describes the issue.
type InputMaxAttemptsError struct {
	message string
}

// NewInputMaxAttemptsError is a constructor that creates and returns a new InputMaxAttemptsError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of InputMaxAttemptsError.
func NewInputMaxAttemptsError(message string) error {
	return &InputMaxAttemptsError{
		message: message,
	}
}

func (err InputMaxAttemptsError) Error() string {
	return err.message
}
