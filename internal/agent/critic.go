package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Critic struct {
	OllamaURL  string
	Model      string
	httpClient *http.Client
}

type CriticEvaluation struct {
	Approved bool   `json:"approved"`
	Feedback string `json:"feedback"`
}

func NewCritic(url, model string) *Critic {
	if url == "" { url = "http://localhost:11434/api/generate" }
	if model == "" { model = "llama3" }
	return &Critic{
		OllamaURL: url,
		Model:     model,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:    10,
				IdleConnTimeout: 90 * time.Second,
			},
		},
	}
}

func (c *Critic) Evaluate(brief *MarketingBrief, personaContent, draft string) (*CriticEvaluation, error) {
	briefJSON, _ := json.MarshalIndent(brief, "", "  ")
	
	prompt := fmt.Sprintf(`### MARKETING MASTERY AUDITOR ###
You are a Zero-Tolerance Marketing Editor and AI Detection Specialist.

### REFERENCE DOCUMENTS ###
STRATEGIC BRIEF:
%s

PERSONA PROFILE:
%s

### CONTENT TO AUDIT ###
%s

### AUDIT CRITERIA ###
1. NO AI FLUFF: Reject any generic filler (e.g., "In today's digital world", "Unlock the power").
2. STRATEGY ALIGNMENT: Does it follow the Framework (%s) and Triggers (%v)?
3. HOOK STRENGTH: Is the first sentence a world-class viral hook?
4. PERSONA INTEGRITY: Does the voice match the profile perfectly?

### OUTPUT INSTRUCTIONS ###
Output in strict JSON format:
{
  "approved": true | false,
  "feedback": "Detailed explanation of flaws or suggestions for improvement"
}`, string(briefJSON), personaContent, draft, brief.Framework, brief.Triggers)

	reqBody := map[string]interface{}{
		"model":  c.Model,
		"prompt": prompt,
		"stream": false,
		"format": "json",
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := c.httpClient.Post(c.OllamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil { return nil, err }
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ollama critic error %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Response string `json:"response"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { return nil, err }

	var eval CriticEvaluation
	if err := json.Unmarshal([]byte(result.Response), &eval); err != nil {
		return nil, fmt.Errorf("failed to parse critic evaluation JSON: %v", err)
	}

	return &eval, nil
}
