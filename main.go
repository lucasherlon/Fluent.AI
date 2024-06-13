package main

import (
	"fluent_ai/gemini"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	arguments := os.Args

	var idiomaSaida string
	if len(arguments) > 1 {
		idiomaSaida = arguments[1]
	} else {
		idiomaSaida = "português brasileiro"
	}

	erro := godotenv.Load(".env")

	if erro != nil {
		fmt.Println("Error loading .env file:", erro)
	}
	gemini.GeneratePrompt(idiomaSaida)
}
