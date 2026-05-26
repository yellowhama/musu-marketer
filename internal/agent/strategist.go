package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Strategist struct {
	OllamaURL string
	Model     string
}

type MarketingBrief struct {
	ValueProp string   `json:"value_proposition"`
	Target    string   `json:"target_segment"`
	Triggers  []string `json:"triggers"`
	Goal      string   `json:"primary_goal"`
	Framework string   `json:"selected_framework"` // New: Guided by Bible
}

func NewStrategist(url, model string) *Strategist {
	if url == "" { url = "http://localhost:11434/api/generate" }
	if model == "" { model = "llama3" }
	return &Strategist{OllamaURL: url, Model: model}
}

func (s *Strategist) CreateBrief(context string) (*MarketingBrief, error) {
	bible := LoadMarketingBible()
	
	prompt := fmt.Sprintf(`### MARKETING MASTERY BIBLE ###
%s

### MISSION ###
You are a Senior Marketing Strategist. Analyze the technical context and select the best professional strategy from the Bible.

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
}`, bible, context)

	reqBody := map[string]interface{}{
		"model":  s.Model,
		"prompt": prompt,
		"stream": false,
		"format": "json",
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := http.Post(s.OllamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil { return nil, err }
	defer resp.Body.Close()

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
