package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/k6zma/cli-hangman/internal/application"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

	jsonPath := flag.String("jsonPath", "words.json", "path to JSON with words")
	flag.Parse()

	gameService := application.NewGameService()

	err := gameService.StartGame(*jsonPath)
	if err != nil {
		os.Exit(0)
	}
}
