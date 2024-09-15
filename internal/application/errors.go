package application

// WordSelectorError defines a custom error for issues during word selection based on difficulty.
//
// Fields:
// - message: the error message that describes the issue.
type WordSelectorError struct {
	message string
}

// NewWordSelectorError is a constructor that creates and returns a new WordSelectorError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of WordSelectorError.
func NewWordSelectorError(message string) error {
	return WordSelectorError{
		message: message,
	}
}

func (err WordSelectorError) Error() string {
	return err.message
}

// WordsLoadingError defines a custom error for issues during words loading.
//
// Fields:
// - message: the error message that describes the issue.
type WordsLoadingError struct {
	message string
}

// NewWordsLoadingError is a constructor that creates and returns a new WordsLoadingError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of WordsLoadingError.
func NewWordsLoadingError(message string) error {
	return WordSelectorError{
		message: message,
	}
}

func (err WordsLoadingError) Error() string {
	return err.message
}

// GamePropertiesCollectingError defines a custom error for issues during game props collecting.
//
// Fields:
// - message: the error message that describes the issue.
type GamePropertiesCollectingError struct {
	message string
}

// NewGamePropertiesCollectingError is a constructor that creates and returns a new GamePropertiesCollectingError.
//
// Parameters:
// - message: a string containing the error message.
//
// Function returns:
// - error: a new instance of GamePropertiesCollectingError.
func NewGamePropertiesCollectingError(message string) error {
	return WordSelectorError{
		message: message,
	}
}

func (err GamePropertiesCollectingError) Error() string {
	return err.message
}
