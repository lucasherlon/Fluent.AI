package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Generate the prompt to be translated
func GeneratePrompt(input, idiomaSaida string) string {

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-latest")
	prompt := fmt.Sprintf("Traduza o seguinte texto para o %s, por favor: %s", idiomaSaida, input)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))

	if err != nil {
		return "Erro ao processar traduçao"
	}
	response := stringfyResponse(resp)
	return response

}

// Show the response generatade from the model
func stringfyResponse(resp *genai.GenerateContentResponse) string {
	var response string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				response += fmt.Sprintf("%s", part)
			}
		}
	}
	return response
}
