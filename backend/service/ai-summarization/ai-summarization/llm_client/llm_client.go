package llm_client

import (
	// "bytes"
	"context"
	// "encoding/json"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/shishir54234/NewsScraper/backend/pkg/config"
	"google.golang.org/genai"
)

type LLMClient interface {
	GenerateDescription(ctx context.Context, content string) (string, error)
}

type openaiClient struct {
	httpClient *http.Client
	apiKey     string
	baseURL    string
}

type geminiAiClient struct {
	client     *genai.Client
	apiKey     string
	baseURL    string
}


func NewLLMClient(llm_client_config *config.LlmConfig) LLMClient {

	config := genai.ClientConfig{
		APIKey:  llm_client_config.ApiKey,
		Backend: genai.BackendGeminiAPI,
	}
	new_client, err:= genai.NewClient(context.Background(), &config)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return &geminiAiClient{
		client:     new_client,
		apiKey:     llm_client_config.ApiKey,
		baseURL:    llm_client_config.BaseUrl,
	}


}

func (g *geminiAiClient) GenerateDescription(ctx context.Context, content string) (string, error) {
	result, err:= g.client.Models.GenerateContent(ctx, "gemini-2.0-flash",
        genai.Text(content),
        nil,
    )
    if err != nil {
        fmt.Println("Error:", err)
		return "", err
    }
    fmt.Println(result.Text())
	return result.Text(), nil
}