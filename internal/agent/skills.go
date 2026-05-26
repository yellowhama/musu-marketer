package agent

import (
	"fmt"
	"os"
)

// LoadMarketingBible reads the high-resolution marketing knowledge base.
func LoadMarketingBible() (string, error) {
	path := "internal/agent/skills/MARKETING_BIBLE.md"
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("marketing bible file not found at %s", path)
	}
	return string(data), nil
}
