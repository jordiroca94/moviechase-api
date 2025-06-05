package recommend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jordiroca94/moviechase-api/utils"
	"github.com/sashabaranov/go-openai"
)

type RecommendPayload struct {
	MovieName string `json:"movie_name"`
}

type RecommendHandler struct {
}

// NewRecommendHandler creates a new instance of RecommendHandler
func NewRecommendHandler() *RecommendHandler {
	return &RecommendHandler{}
}

func (h *RecommendHandler) handleGetRecommendation(w http.ResponseWriter, r *http.Request) {

	var payload RecommendPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	movieName := payload.MovieName

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or failed to load:", err)
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("OPENAI_API_KEY environment variable not set")
	}

	tplBytes, err := os.ReadFile("service/recommend/prompts/prompt.tpl")
	if err != nil {
		log.Fatalf("Failed to read prompt: %v\n", err)
	}

	tpl, err := template.New("prompt").Parse(string(tplBytes))
	if err != nil {
		log.Fatalf("Failed to parse prompt: %v\n", err)
	}

	var renderedPrompt bytes.Buffer
	err = tpl.Execute(&renderedPrompt, struct {
		MovieName string
	}{
		MovieName: movieName,
	})
	if err != nil {
		log.Fatalf("Failed to execute prompt template: %v\n", err)
	}

	client := openai.NewClient(apiKey)
	ctx := context.Background()

	chatReq := openai.ChatCompletionRequest{
		Model: "gpt-4o",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: renderedPrompt.String(),
			},
		},
		MaxTokens: 1000,
	}

	stop := utils.StartTimer()

	resp, err := client.CreateChatCompletion(ctx, chatReq)
	stop()

	fmt.Println()
	if err != nil {
		log.Fatalf("OpenAI API error: %v\n", err)
	}

	if len(resp.Choices) > 0 {
		utils.WriteJSON(w, http.StatusOK, map[string]string{"response": resp.Choices[0].Message.Content})

	} else {
		http.Error(w, "No recommendations found", http.StatusNotFound)
		return
	}

}
