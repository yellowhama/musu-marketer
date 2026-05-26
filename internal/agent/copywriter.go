package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Copywriter struct {
	OllamaURL      string
	Model          string
	ActivePersona  string
	PersonaContent string
	ProjectPath    string
	httpClient     *http.Client // Optimized: Reusable client pool
}

func NewCopywriter(url, model, personaName, projectPath string) *Copywriter {
	if url == "" { url = "http://localhost:11434/api/generate" }
	if model == "" { model = "llama3" }
	
	c := &Copywriter{
		OllamaURL:     url, 
		Model:         model, 
		ActivePersona: personaName, 
		ProjectPath:   projectPath,
		httpClient: &http.Client{
			Timeout: 120 * time.Second, // Drafting takes longer
			Transport: &http.Transport{
				MaxIdleConns:        50,
				IdleConnTimeout:     90 * time.Second,
				MaxIdleConnsPerHost: 10,
			},
		},
	}
	c.loadPersona(personaName)
	return c
}

func (c *Copywriter) loadPersona(name string) {
	path := filepath.Join(c.ProjectPath, "personas", name+".md")
	data, err := os.ReadFile(path)
	if err != nil {
		path = filepath.Join("personas", name+".md")
		data, err = os.ReadFile(path)
		if err != nil {
			c.PersonaContent = "Role: Professional Marketer\nTone: Neutral"
			return
		}
	}
	c.PersonaContent = string(data)
}

func (c *Copywriter) GenerateCampaign(topic string, context string, brief *MarketingBrief) (string, error) {
	bible, err := LoadMarketingBible()
	if err != nil {
		return "", fmt.Errorf("critical: cannot draft without marketing bible: %v", err)
	}
	
	briefJSON, _ := json.MarshalIndent(brief, "", "  ")
	
	systemPrompt := fmt.Sprintf(`### MARKETING MASTERY BIBLE ###
%s

### ACTIVE PERSONA PROFILE ###
%s

### STRATEGIC BRIEF ###
%s

### MISSION ###
Execute a high-impact marketing campaign based on the Context and Strategic Brief.
You MUST strictly use the formulas and patterns defined in the MARKETING MASTERY BIBLE.

Tasks:
1. Apply the Selected Framework (%s) flawlessly.
2. Use the specified Triggers to drive emotional resonance.
3. Use a 2024 Viral Hook Pattern for the opening.

Context:
%s

Requirements:
- Output a high-impact Twitter Thread (5-7 posts).
- Output a professional Blog Post / LinkedIn Article.`, bible, c.PersonaContent, string(briefJSON), brief.Framework, context)

	reqBody := map[string]interface{}{
		"model":  c.Model,
		"prompt": systemPrompt,
		"stream": false,
	}

	jsonData, _ := json.Marshal(reqBody)
	resp, err := c.httpClient.Post(c.OllamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama copywriter error %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Response string `json:"response"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Response, nil
}
