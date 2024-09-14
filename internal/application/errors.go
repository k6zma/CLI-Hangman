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
