package infrastructure

import (
	"unicode"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/utils"

	"fmt"
	"os"
	"time"
)

// printTextPerSymbol prints the given text character by character with a delay.
//
// Parameters:
// - text: the string to be printed.
// - delay: the duration to wait between printing each character.
func printTextPerSymbol(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
}

// printASCIIArt prints ASCII art for the Hangman game and CLI game title.
func printASCIIArt() {
	hangmanASCII := `
 _   _    _    _   _  ____ __  __    _    _   _ 
| | | |  / \  | \ | |/ ___|  \/  |  / \  | \ | |
| |_| | / _ \ |  \| | |  _| |\/| | / _ \ |  \| |
|  _  |/ ___ \| |\  | |_| | |  | |/ ___ \| |\  |
|_| |_/_/   \_\_| \_|\____|_|  |_/_/   \_\_| \_|
`
	cliGameASCII := `
  ____ _     ___      ____    _    __  __ _____ 
 / ___| |   |_ _|    / ___|  / \  |  \/  | ____|
| |   | |    | |    | |  _  / _ \ | |\/| |  _|  
| |___| |___ | |    | |_| |/ ___ \| |  | | |___ 
 \____|_____|___|    \____/_/   \_\_|  |_|_____|
`
	delayOfPrintingSymbols := 3 * time.Millisecond

	printTextPerSymbol(hangmanASCII, delayOfPrintingSymbols)
	printTextPerSymbol(cliGameASCII, delayOfPrintingSymbols)

	time.Sleep(1 * time.Second)
}

// printRules prints the game rules to the terminal.
func printRules() {
	fmt.Println("Here you can read the rules: ")
	fmt.Println("1) Available languages - Russian and English")
	fmt.Println("2) Available difficulties - easy, medium and hard")
	fmt.Println("3) The number of attempts must be in the range from 1 to 99")
	fmt.Println("4) You can only enter one letter at a time")
	fmt.Println("5) If you enter the wrong letter, you lose one attempt")
	fmt.Println("6) If you have less than 20% of attempts left, you will be offered a hint")
	fmt.Println("")
	fmt.Println("Now you can agree with the rules or not, type agree or disagree")
}

// HangmanIntro displays the introduction for the Hangman game and handles user agreement to the rules.
func HangmanIntro() {
	utils.ClearScreen()
	printASCIIArt()

	utils.ClearScreen()
	fmt.Println("-------------------------------------------")
	fmt.Println("       WELCOME TO HANGMAN - CLI GAME       ")
	fmt.Println("-------------------------------------------")
	fmt.Println("")

	printRules()

	for {
		flagSuggestion, err := AcceptTheRules()

		if err == nil {
			if flagSuggestion {
				fmt.Println("You agreed to the rules!")
			} else {
				fmt.Println("You disagreed with the rules. So, bye")
				os.Exit(0)
			}

			break
		}

		fmt.Println("Error:", err)
	}
}

// PrintWordState prints the current state of the word being guessed.
//
// Parameters:
// - game: a pointer to the Game object containing the game state.
func PrintWordState(game *domain.Game) {
	wordLetters, err := game.GetWordLetters()
	if err != nil {
		fmt.Println("Error getting word letters:", err)
		return
	}

	guessedLetters, err := game.GetGuessedLetters()
	if err != nil {
		fmt.Println("Error getting guessed letters:", err)
		return
	}

	wordState := ""

	for _, letter := range wordLetters {
		if guessedLetters[letter] {
			wordState += string(letter) + " "
		} else {
			wordState += "_ "
		}
	}

	fmt.Println("Current word state:", wordState)
}

// PrintAvailableLetters prints the letters that have not been guessed yet.
//
// Parameters:
// - game: a pointer to the Game object containing the game state.
func PrintAvailableLetters(game *domain.Game) {
	var alphabet []rune

	language, err := game.GetLanguage()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if language == "en" {
		alphabet = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	} else {
		alphabet = []rune("АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ")
	}

	availableLetters := ""

	guessedLetters, err := game.GetGuessedLetters()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, letter := range alphabet {
		lowerLetter := unicode.ToLower(letter)

		if !guessedLetters[lowerLetter] && !guessedLetters[unicode.ToUpper(lowerLetter)] {
			availableLetters += string(letter) + " "
		}
	}

	fmt.Println("Available letters:", availableLetters)
}

// PrintHangman prints the current hangman stage based on the number of attempts left.
//
// Parameters:
// - game: a pointer to the Game object containing the game state.
func PrintHangman(game *domain.Game) {
	stages := []string{
		`
  +---+
      |
      |
      |
     ===`, `
  +---+
  O   |
      |
      |
     ===`, `
  +---+
  O   |
  |   |
      |
     ===`, `
  +---+
  O   |
 /|   |
      |
     ===`, `
  +---+
  O   |
 /|\  |
      |
     ===`, `
  +---+
  O   |
 /|\  |
 /    |
     ===`, `
  +---+
  O   |
 /|\  |
 / \  |
     ===`,
	}

	currentAttempts, err := game.GetCurrentAttempts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	maxAttempts, err := game.GetMaxAttempts()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	percentage := float64(currentAttempts) / float64(maxAttempts)
	stageIndex := int(percentage * float64(len(stages)-1))

	if stageIndex >= len(stages) {
		fmt.Println(stages[len(stages)-1])
	} else {
		fmt.Println(stages[stageIndex])
	}
}

// PrintHint prints a hint if the player has less than 30% of attempts left.
//
// Parameters:
// - game: a pointer to the Game object containing the game state.
func PrintHint(game *domain.Game) {
	currentAttempts, err := game.GetCurrentAttempts()
	if err != nil {
		fmt.Println("Error getting current attempts:", err)
	}

	maxAttempts, err := game.GetMaxAttempts()
	if err != nil {
		fmt.Println("Error getting max attempts:", err)
	}

	if float64(currentAttempts)/float64(maxAttempts) >= 0.8 {
		hint, _ := game.GetWordHint()
		fmt.Println("Hint:", hint)
	}
}

// PrintGameOver prints the GameOver ASCII art.
func PrintGameOver() {
	gameASCII := `
  ____    _    __  __ _____ 
 / ___|  / \  |  \/  | ____|
| |  _  / _ \ | |\/| |  _|  
| |_| |/ ___ \| |  | | |___ 
 \____/_/   \_\_|  |_|_____|
`
	overASCII := `
  _____     _______ ____  
 / _ \ \   / / ____|  _ \ 
| | | \ \ / /|  _| | |_) |
| |_| |\ V / | |___|  _ < 
 \___/  \_/  |_____|_| \_\
`
	delayOfPrintingSymbols := 2 * time.Millisecond

	printTextPerSymbol(gameASCII, delayOfPrintingSymbols)
	printTextPerSymbol(overASCII, delayOfPrintingSymbols)
}

// PrintVictory prints the win ASCII art.
func PrintVictory() {
	victoryASCII := `
__   _____  _   _  __        _____ _   _ _ 
\ \ / / _ \| | | | \ \      / /_ _| \ | | |
 \ V / | | | | | |  \ \ /\ / / | ||  \| | |
  | || |_| | |_| |   \ V  V /  | || |\  |_|
  |_| \___/ \___/     \_/\_/  |___|_| \_(_)
`
	delayOfPrintingSymbols := 2 * time.Millisecond

	printTextPerSymbol(victoryASCII, delayOfPrintingSymbols)
}
