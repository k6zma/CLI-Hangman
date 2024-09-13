package domain

type WordLoaderError struct {
	message string
}

func NewWordLoaderError(message string) error {
	return &WordLoaderError{
		message: message,
	}
}

func (err WordLoaderError) Error() string {
	return err.message
}
