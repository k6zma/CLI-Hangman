package application

import (
	"fmt"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/utils"
)

type GameService struct{}

// NewGameService creates and returns a new GameService object.
func NewGameService() *GameService {
	return &GameService{}
}

// LoadWords loads words from the given JSON file.
//
// Parameters:
// - wordsFilename: the path to the JSON file containing the words.
//
// Function returns:
// - *domain.ParsedWords: a pointer to the structure containing the parsed words.
// - error: an error if the file cannot be read or parsed.
func (s *GameService) LoadWords(wordsFilename string) (*domain.ParsedWords, error) {
	wordsData, err := infrastructure.LoadWords(wordsFilename)
	if err != nil {
		return nil, NewWordsLoadingError("error while loading the words")
	}

	return wordsData, nil
}

// InitializeGameProperties initializes game properties by displaying an introduction and requesting user input.
//
// Function returns:
// - *domain.GameProperties: the initialized game properties.
func (s *GameService) InitializeGameProperties() *domain.GameProperties {
	infrastructure.HangmanIntro()

	gameProperties := infrastructure.RequestGameProperties()

	return gameProperties
}

// SelectWordByProperties selects a word based on the provided game properties.
//
// Parameters:
// - wD: the parsed words data containing words in different languages.
// - gP: the game properties which include language and difficulty.
//
// Function returns:
// - *domain.WordWithHintJSON: a pointer to the selected word with a hint.
// - error: an error if there is a failure to retrieve language, difficulty, or to select a word.
func (s *GameService) SelectWordByProperties(wD *domain.ParsedWords, gP *domain.GameProperties) (*domain.WordWithHintJSON, error) {
	language, err := gP.GetLanguageFromProperties()
	if err != nil {
		return nil, err
	}

	difficulty, err := gP.GetDifficultyFromProperties()
	if err != nil {
		return nil, err
	}

	var selectedWords []domain.WordWithHintJSON

	switch *language {
	case "en":
		selectedWords = wD.EnWords
	case "ru":
		selectedWords = wD.RuWords
	}

	selectedWord, err := SelectWordByDifficulty(selectedWords, *difficulty)
	if err != nil {
		return nil, err
	}

	return selectedWord, nil
}

// StartGameSession initializes a new game session with the selected word and game properties.
//
// Parameters:
// - selectedWord: the word selected based on language and difficulty.
// - gameProperties: the properties containing language, difficulty, and the maximum number of attempts.
//
// Function returns:
// - *domain.Game: a pointer to the initialized game.
// - error: an error if any game property fails to be retrieved.
func (s *GameService) StartGameSession(selectedWord *domain.WordWithHintJSON, gameProperties *domain.GameProperties) (*domain.Game, error) {
	language, err := gameProperties.GetLanguageFromProperties()
	if err != nil {
		return nil, NewGamePropertiesCollectingError("failed to get language")
	}

	difficulty, err := gameProperties.GetDifficultyFromProperties()
	if err != nil {
		return nil, NewGamePropertiesCollectingError("failed to get difficulty")
	}

	maxAttempts, err := gameProperties.GetMaxAttemptsFromProperties()
	if err != nil {
		return nil, NewGamePropertiesCollectingError("failed to get max attempts")
	}

	word := domain.NewWord(*selectedWord, *language, *difficulty)
	game := domain.NewGame(word, *maxAttempts)

	return game, nil
}

// RunGameLoop executes the main game loop where the player guesses letters.
//
// Parameters:
// - game: the pointer to the game object containing the current game state.
//
// Function returns:
// - error: an error if the game encounters an issue while retrieving word letters at the end.
func (s *GameService) RunGameLoop(game *domain.Game) error {
	for !game.IsGameOver() && !game.IsGameWon() {
		utils.ClearScreen()

		infrastructure.PrintHangman(game)
		infrastructure.PrintWordState(game)
		infrastructure.PrintAvailableLetters(game)
		infrastructure.PrintHint(game)

		letter, err := infrastructure.GetLetterFromUser()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if game.GuessLetter(letter) {
			fmt.Println("Correct letter guess!")
		} else {
			fmt.Println("Wrong letter guess!")
		}
	}

	utils.ClearScreen()

	infrastructure.PrintHangman(game)
	infrastructure.PrintWordState(game)

	if game.IsGameWon() {
		infrastructure.PrintVictory()
	} else {
		wordLetters, err := game.GetWordLetters()
		if err != nil {
			return err
		}

		infrastructure.PrintGameOver()
		fmt.Println("\nThe word was:", string(wordLetters))
	}

	return nil
}

// StartGame loads the words, initializes the game properties, selects a word, and runs the game loop.
//
// Parameters:
// - wordsFilename: the path to the JSON file containing the word list.
//
// Function returns:
// - error: an error if there is an issue at any stage of the game.
func (s *GameService) StartGame(wordsFilename string) error {
	wordsData, err := s.LoadWords(wordsFilename)
	if err != nil {
		return err
	}

	gameProperties := s.InitializeGameProperties()

	selectedWord, err := s.SelectWordByProperties(wordsData, gameProperties)
	if err != nil {
		return err
	}

	game, err := s.StartGameSession(selectedWord, gameProperties)
	if err != nil {
		return err
	}

	return s.RunGameLoop(game)
}
