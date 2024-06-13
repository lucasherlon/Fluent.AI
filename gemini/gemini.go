package gemini

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Generate the prompt to be translated
func GeneratePrompt(idiomaSaida string) {

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-latest")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(">> Insira o texto que você deseja traduzir: ")

		input, _ := reader.ReadString('\n')
		prompt := fmt.Sprintf("Por gentileza, traduza o seguinte texto para o %v: %v", idiomaSaida, input)

		resp, err := model.GenerateContent(ctx, genai.Text(prompt))

		if err != nil {
			log.Fatal(err)
		}
		printResponse(resp)
	}
}

// Show the response generatad from the model
func printResponse(resp *genai.GenerateContentResponse) {
	fmt.Println("")
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("")
}
