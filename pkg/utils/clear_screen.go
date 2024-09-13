package utils

import "fmt"

// ClearScreen clears the terminal screen using escape sequences
func ClearScreen() {
	fmt.Print("\"\\033[H\\033[2J\"")
}
