package main

import (
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/application"
)

func main() {
	wordsFilename := "words.json"

	gameService := application.NewGameService()

	err := gameService.StartGame(wordsFilename)
	if err != nil {
		os.Exit(0)
	}
}
