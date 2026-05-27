package bridge

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type WikiBridge struct {
	WikiDir string
}

type KnowledgeSource struct {
	ID      string
	Title   string
	Content string
	Source  string
	Project string
}

type indexEntry struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Source  string   `json:"source"`
	Project string   `json:"project"`
	Path    string   `json:"path"`
	Tags    []string `json:"tags,omitempty"`
	Summary string   `json:"summary,omitempty"`
}

func NewWikiBridge(wikiDir string) *WikiBridge {
	return &WikiBridge{WikiDir: wikiDir}
}

func (b *WikiBridge) FindByTopic(topic string) ([]KnowledgeSource, error) {
	if strings.TrimSpace(topic) == "" {
		return nil, nil
	}
	if results, err := b.findByTopicFromIndex(topic); err == nil && len(results) > 0 {
		return results, nil
	}
	return b.findByTopicFromFiles(topic)
}

func (b *WikiBridge) findByTopicFromIndex(topic string) ([]KnowledgeSource, error) {
	indexPath := filepath.Join(b.WikiDir, "index.json")
	data, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, err
	}

	var entries []indexEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}

	query := strings.ToLower(strings.TrimSpace(topic))
	var results []KnowledgeSource
	seen := map[string]bool{}
	for _, entry := range entries {
		path := filepath.Join(b.WikiDir, entry.Path)
		body := ""
		if raw, readErr := os.ReadFile(path); readErr == nil {
			body = string(raw)
		}
		if !matchesTopic(query, entry, body) {
			continue
		}
		if seen[path] {
			continue
		}
		seen[path] = true
		results = append(results, KnowledgeSource{
			ID:      entry.ID,
			Title:   entry.Title,
			Content: body,
			Source:  entry.Source,
			Project: entry.Project,
		})
	}
	return results, nil
}

func (b *WikiBridge) findByTopicFromFiles(topic string) ([]KnowledgeSource, error) {
	var results []KnowledgeSource
	query := strings.ToLower(strings.TrimSpace(topic))
	
	err := filepath.Walk(b.WikiDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		body := string(data)
		if strings.Contains(strings.ToLower(filepath.Base(path)), query) || strings.Contains(strings.ToLower(body), query) {
			results = append(results, KnowledgeSource{
				ID:      filepath.Base(path),
				Title:   filepath.Base(path),
				Content: body,
			})
		}
		return nil
	})
	
	return results, err
}

func matchesTopic(query string, entry indexEntry, body string) bool {
	if strings.Contains(strings.ToLower(entry.Title), query) {
		return true
	}
	if strings.Contains(strings.ToLower(entry.Summary), query) {
		return true
	}
	if strings.Contains(strings.ToLower(entry.Path), query) {
		return true
	}
	for _, tag := range entry.Tags {
		if strings.Contains(strings.ToLower(tag), query) {
			return true
		}
	}
	if strings.Contains(strings.ToLower(body), query) {
		return true
	}
	return false
}
