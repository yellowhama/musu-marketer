package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Strategist struct {
	OllamaURL  string
	Model      string
	httpClient *http.Client // Optimized: Reusable client pool
}

type MarketingBrief struct {
	ValueProp string   `json:"value_proposition"`
	Target    string   `json:"target_segment"`
	Triggers  []string `json:"triggers"`
	Goal      string   `json:"primary_goal"`
	Framework string   `json:"selected_framework"`
}

func NewStrategist(url, model string) *Strategist {
	if url == "" { url = "http://localhost:11434/api/generate" }
	if model == "" { model = "llama3" }
	return &Strategist{
		OllamaURL: url,
		Model:     model,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        50,
				IdleConnTimeout:     90 * time.Second,
				MaxIdleConnsPerHost: 10,
			},
		},
	}
}

func (s *Strategist) CreateBrief(context string, history string) (*MarketingBrief, error) {
	bible, err := LoadMarketingBible()
	if err != nil {
		return nil, fmt.Errorf("critical: strategy cannot proceed without marketing bible: %v", err)
	}
	
	historySection := ""
	if history != "" {
		historySection = fmt.Sprintf("\n### PREVIOUS SUCCESSFUL CAMPAIGNS ###\n%s\n\nEnsure this new strategy builds upon the above history and does not repeat the exact same angles.", history)
	}

	prompt := fmt.Sprintf(`### MARKETING MASTERY BIBLE ###
%s

### MISSION ###
You are a Senior Marketing Strategist. Analyze the technical context and select the best professional strategy from the Bible.
%s

Context:
%s

Tasks:
1. Select the most effective Framework (PAS, AIDA, or BAB) from the Bible.
2. Select 2 Psychological Triggers from the Bible.
3. Define the Segment and Goal.

Output in strict JSON format:
{
  "value_proposition": "...",
  "target_segment": "...",
  "triggers": ["...", "..."],
  "primary_goal": "...",
  "selected_framework": "AIDA | PAS | BAB"
}`, bible, historySection, context)

	reqBody := map[string]interface{}{
		"model":  s.Model,
		"prompt": prompt,
		"stream": false,
		"format": "json",
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := s.httpClient.Post(s.OllamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil { return nil, err }
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama strategist error %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Response string `json:"response"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { return nil, err }

	var brief MarketingBrief
	if err := json.Unmarshal([]byte(result.Response), &brief); err != nil {
		return nil, fmt.Errorf("failed to parse brief JSON: %v", err)
	}

	return &brief, nil
}
