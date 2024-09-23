package application

// Сonstants with error texts.
const (
	ErrWordSelection            = "error during word selection based on difficulty"
	ErrWordsLoading             = "error during words loading"
	ErrGamePropertiesCollecting = "error during game properties getting"
)

// WordSelectorError defines a custom error for issues during word selection based on difficulty.
//
// Fields:
// - message: the error message that describes the issue.
type WordSelectorError struct {
	message string
}

// NewWordSelectorError is a constructor that creates and returns a new WordSelectorError.
//
// Function returns:
// - error: a new instance of WordSelectorError.
func NewWordSelectorError() error {
	return &WordSelectorError{
		message: ErrWordSelection,
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
// Function returns:
// - error: a new instance of WordsLoadingError.
func NewWordsLoadingError() error {
	return &WordsLoadingError{
		message: ErrWordsLoading,
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
// Function returns:
// - error: a new instance of GamePropertiesCollectingError.
func NewGamePropertiesCollectingError() error {
	return &GamePropertiesCollectingError{
		message: ErrGamePropertiesCollecting,
	}
}

func (err GamePropertiesCollectingError) Error() string {
	return err.message
}
