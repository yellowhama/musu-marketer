package bridge

import (
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

func NewWikiBridge(wikiDir string) *WikiBridge {
	return &WikiBridge{WikiDir: wikiDir}
}

func (b *WikiBridge) FindByTopic(topic string) ([]KnowledgeSource, error) {
	var results []KnowledgeSource
	
	err := filepath.Walk(b.WikiDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		
		if strings.Contains(strings.ToLower(filepath.Base(path)), strings.ToLower(topic)) {
			data, err := os.ReadFile(path)
			if err == nil {
				results = append(results, KnowledgeSource{
					ID:      filepath.Base(path),
					Title:   filepath.Base(path),
					Content: string(data),
				})
			}
		}
		return nil
	})
	
	return results, err
}
