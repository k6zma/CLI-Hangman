package infrastructure

// Сonstants with error texts.
const (
	ErrWordLoading          = "error loading words"
	ErrLanguageInput        = "invalid language input"
	ErrDifficultyInput      = "invalid difficulty input"
	ErrInputMaxAttempts     = "invalid max attempts input"
	ErrInputRulesSuggestion = "invalid rules suggestion input"
	ErrLetterInput          = "invalid letter input"
)

// WordLoaderError defines a custom error of word loading.
//
// Fields:
// - message: the error message that describes the issue.
type WordLoaderError struct {
	message string
}

// NewWordLoaderError is constructor which creates and returns a new WordLoaderError.
//
// Function returns:
// - error: a new instance of WordLoaderError.
func NewWordLoaderError() error {
	return &WordLoaderError{
		message: ErrWordLoading,
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
// Function returns:
// - error: a new instance of InputLanguageError.
func NewInputLanguageError() error {
	return &InputLanguageError{
		message: ErrLanguageInput,
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
// Function returns:
// - error: a new instance of InputDifficultyError.
func NewInputDifficultyError() error {
	return &InputDifficultyError{
		message: ErrDifficultyInput,
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
// Function returns:
// - error: a new instance of InputMaxAttemptsError.
func NewInputMaxAttemptsError() error {
	return &InputMaxAttemptsError{
		message: ErrInputMaxAttempts,
	}
}

func (err InputMaxAttemptsError) Error() string {
	return err.message
}

// InputRulesSuggestionError defines a custom error for invalid rules suggestion input.
//
// Fields:
// - message: the error message that describes the issue.
type InputRulesSuggestionError struct {
	message string
}

// NewInputRulesSuggestionError is a constructor that creates and returns a new InputRulesSuggestionError.
//
// Function returns:
// - error: a new instance of InputRulesSuggestionError.
func NewInputRulesSuggestionError() error {
	return &InputRulesSuggestionError{
		message: ErrInputRulesSuggestion,
	}
}

func (err InputRulesSuggestionError) Error() string {
	return err.message
}

// InputLetterError defines a custom error for invalid letter input.
//
// Fields:
// - message: the error message that describes the issue.
type InputLetterError struct {
	message string
}

// NewInputLetterError is a constructor that creates and returns a new InputLetterError.
//
// Function returns:
// - error: a new instance of InputLetterError.
func NewInputLetterError() error {
	return &InputLetterError{
		message: ErrLetterInput,
	}
}

func (err InputLetterError) Error() string {
	return err.message
}
