package infrastructure

import (
	"fmt"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/utils"
	"os"
	"time"
)

func printTextPerSymbol(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(delay)
	}
}

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

func printRules() {
	fmt.Println("Here you can read the rules: ")
	fmt.Println("1) Available languages - Russian and English")
	fmt.Println("2) Available difficulties - easy, medium and hard")
	fmt.Println("3) The number of attempts must be in the range from 1 to 99")
	fmt.Println("4) You can only enter one letter at a time")
	fmt.Println("5) If you enter the wrong letter, you lose one attempt")
	fmt.Println("6) If you have less than 30% of attempts left, you will be offered a hint")
	fmt.Println("")
	fmt.Println("Now you can agree with the rules or not, type agree or disagree")
}

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
