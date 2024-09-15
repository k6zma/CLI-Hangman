package infrastructure

import (
	"bufio"
	"os"
	"unicode/utf8"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/utils"

	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// inputLanguage collect the user input the game language.
//
// Function returns:
// - string: the selected language.
// - error: an error if an invalid language was entered.
func inputLanguage() (string, error) {
	var language string
	if _, err := fmt.Scanln(&language); err != nil {
		return "", NewInputLanguageError("error with scan line")
	}

	language = strings.TrimSpace(language)

	switch language {
	case "russian", "ru", "ру", "русский":
		return "ru", nil
	case "english", "en", "англ", "английский":
		return "en", nil
	default:
		return "", NewInputLanguageError("wrong language entered")
	}
}

// inputDifficulty collect the user input the game difficulty.
//
// Function returns:
// - string: the selected difficulty.
// - error: an error if an invalid difficulty was entered.
func inputDifficulty() (string, error) {
	var difficulty string
	if _, err := fmt.Scanln(&difficulty); err != nil {
		if err.Error() == "unexpected newline" {
			difficulty = ""
		} else {
			return "", NewInputDifficultyError("error with scan line")
		}
	}

	difficulty = strings.TrimSpace(difficulty)

	switch difficulty {
	case "easy", "изи", "легкий":
		return "easy", nil
	case "medium", "средний":
		return "medium", nil
	case "hard", "хард", "сложный":
		return "hard", nil
	case "":
		difficulties := []string{"easy", "medium", "hard"}

		maxIndex := big.NewInt(int64(len(difficulties)))
		randomIndex, err := rand.Int(rand.Reader, maxIndex)

		if err != nil {
			return "", NewInputDifficultyError("error generating random difficulty")
		}

		randomDifficulty := difficulties[randomIndex.Int64()]

		return randomDifficulty, nil
	default:
		return "", NewInputDifficultyError("wrong difficulty entered")
	}
}

// inputMaxAttempts collect the user input the maximum number of attempts.
//
// Function returns:
// - int: the number of maximum attempts.
// - error: an error if an invalid number of attempts was entered.
func inputMaxAttempts() (int, error) {
	var attemptsStr string
	if _, err := fmt.Scanln(&attemptsStr); err != nil {
		return -1, NewInputMaxAttemptsError("error with scan line")
	}

	maxAttempts, _ := strconv.Atoi(attemptsStr)

	if maxAttempts <= 0 {
		return -1, NewInputMaxAttemptsError("the number of maximum attempts must be greater than 0")
	} else if maxAttempts >= 100 {
		return -1, NewInputMaxAttemptsError("the number of maximum attempts must be less than 100")
	}

	return maxAttempts, nil
}

// RequestGameProperties collect the user input all the necessary game properties.
//
// Function returns:
// - *domain.GameProperties: the game properties entered by the user.
func RequestGameProperties() *domain.GameProperties {
	var (
		language    string
		difficulty  string
		maxAttempts int
		err         error
	)

	utils.ClearScreen()

	for {
		fmt.Println("Please enter your language (en|ru):")

		language, err = inputLanguage()

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		break
	}

	utils.ClearScreen()

	for {
		fmt.Println("Please enter difficulty (easy|medium|hard):")

		difficulty, err = inputDifficulty()

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		break
	}

	utils.ClearScreen()

	for {
		fmt.Println("Please enter the number of attempts (more than 0, but less than 100):")

		maxAttempts, err = inputMaxAttempts()

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		break
	}

	return domain.NewGameProperties(language, difficulty, maxAttempts)
}

// AcceptTheRules collects user input to accept or reject the rules.
//
// Function returns:
// - bool: true if the user agrees, false otherwise.
// - error: an error if the input is invalid or an unknown word is entered.
func AcceptTheRules() (bool, error) {
	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return false, NewInputRulesSuggestionError("error with input suggestion")
	}

	switch input {
	case "agree":
		return true, nil
	case "disagree":
		return false, nil
	default:
		return false, NewInputRulesSuggestionError("you entered an unknown word")
	}
}

// isLetter checks if the given rune is a letter.
//
// Parameters:
// - letter: a symbol represented in rune.
//
// Function returns:
// - bool: true if the rune is a letter, false otherwise.
func isLetter(letter rune) bool {
	if (letter >= 'A' && letter <= 'Z') || (letter >= 'a' && letter <= 'z') {
		return true
	}

	if (letter >= 'А' && letter <= 'Я') || (letter >= 'а' && letter <= 'я') {
		return true
	}

	return false
}

// GetLetterFromUser collects a single letter input from the user.
//
// Function returns:
// - rune: the letter entered by the user.
// - error: an error if the input is invalid or not a single letter.
func GetLetterFromUser() (rune, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter a letter:")

	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	if utf8.RuneCountInString(input) != 1 {
		return 0, NewInputLetterError("please enter only one letter")
	}

	letter, _ := utf8.DecodeRuneInString(input)
	if !isLetter(letter) {
		return 0, NewInputLetterError("please enter a valid letter")
	}

	return letter, nil
}
