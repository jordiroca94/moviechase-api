package recommend

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

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

	apiKey := strings.TrimSpace(os.Getenv("OPENAI_API_KEY"))
	if apiKey == "" {
		log.Println("OPENAI_API_KEY not set")
		http.Error(w, "Server configuration error", http.StatusInternalServerError)
		return
	}

	tplBytes, err := os.ReadFile("service/recommend/prompts/prompt.tpl")
	if err != nil {
		log.Printf("Failed to read prompt: %v\n", err)
		http.Error(w, "Template read error", http.StatusInternalServerError)
		return
	}

	tpl, err := template.New("prompt").Parse(string(tplBytes))
	if err != nil {
		log.Printf("Failed to parse template: %v\n", err)
		http.Error(w, "Template parse error", http.StatusInternalServerError)
		return
	}

	var renderedPrompt bytes.Buffer
	err = tpl.Execute(&renderedPrompt, struct{ MovieName string }{MovieName: payload.MovieName})
	if err != nil {
		log.Printf("Template execution failed: %v\n", err)
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	// OpenAI call
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

	log.Printf("OPENAI_API_KEY starts with: %s", apiKey[:5])
	log.Printf("Rendered prompt: %s", renderedPrompt.String())
	log.Printf("Calling OpenAI with model: %s", chatReq.Model)

	var apiErr *openai.APIError
	if errors.As(err, &apiErr) {
		log.Printf("Status: %d | Type: %s | Message: %s\n", apiErr.HTTPStatusCode, apiErr.Type, apiErr.Message)
	}

	resp, err := client.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		log.Printf("OpenAI API error: %v\n", err)
		http.Error(w, "OpenAI API error", http.StatusInternalServerError)
		return
	}

	if len(resp.Choices) == 0 {
		http.Error(w, "No recommendations found", http.StatusNotFound)
		return
	}

	content := resp.Choices[0].Message.Content
	cleaned := utils.StripCodeBlock(content)

	var recommendations []map[string]interface{}
	if err := json.Unmarshal([]byte(cleaned), &recommendations); err != nil {
		http.Error(w, "Failed to parse recommendations JSON", http.StatusInternalServerError)
		log.Printf("JSON parse error: %v\nOriginal content:\n%s", err, cleaned)
		return
	}

	utils.WriteJSON(w, http.StatusOK, recommendations)

}
