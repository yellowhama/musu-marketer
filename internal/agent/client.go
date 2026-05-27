package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// AgentClient is a generic OpenAI-compatible LLM client.
type AgentClient struct {
	BaseURL     string
	Model       string
	httpClient  *http.Client
}

func NewAgentClient(baseURL, model string) *AgentClient {
	if baseURL == "" {
		baseURL = "http://localhost:11434/v1" // Default to local Ollama OpenAI endpoint
	}
	if model == "" {
		model = "llama3"
	}

	return &AgentClient{
		BaseURL: baseURL,
		Model:   model,
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				IdleConnTimeout:     90 * time.Second,
				MaxIdleConnsPerHost: 20,
			},
		},
	}
}

// OpenAI Chat Completion structures
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type ChatRequest struct {
	Model          string          `json:"model"`
	Messages       []ChatMessage   `json:"messages"`
	Stream         bool            `json:"stream"`
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatMessage `json:"message"`
	} `json:"choices"`
}

// Ask sends a chat completion request.
func (c *AgentClient) Ask(prompt string, jsonFormat bool) (string, error) {
	reqBody := ChatRequest{
		Model: c.Model,
		Messages: []ChatMessage{
			{Role: "user", Content: prompt},
		},
		Stream: false,
	}

	if jsonFormat {
		reqBody.ResponseFormat = &ResponseFormat{Type: "json_object"}
	}

	jsonData, _ := json.Marshal(reqBody)
	url := c.BaseURL + "/chat/completions"

	resp, err := c.httpClient.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("AI connection failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("AI returned error %d: %s", resp.StatusCode, string(body))
	}

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", fmt.Errorf("failed to decode AI response: %v", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("AI returned no choices")
	}

	return chatResp.Choices[0].Message.Content, nil
}
