package domain

// WordLoaderError defines a custom error of word loading
//
// Fields:
// - message: the error message that describes the issue
type WordLoaderError struct {
	message string
}

// NewWordLoaderError is constructor which creates and returns a new WordLoaderError
//
// Parameters:
// - message: a string containing the error message
//
// Function returns:
// - error: a new instance of WordLoaderError
func NewWordLoaderError(message string) error {
	return &WordLoaderError{
		message: message,
	}
}

func (err WordLoaderError) Error() string {
	return err.message
}
