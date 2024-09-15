package main

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/application"

	"os"
)

func main() {
	wordsFilename := "words.json"

	gameService := application.NewGameService()

	err := gameService.StartGame(wordsFilename)
	if err != nil {
		os.Exit(0)
	}
}
