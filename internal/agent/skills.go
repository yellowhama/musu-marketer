package agent

import (
	"os"
)

// LoadMarketingBible reads the high-resolution marketing knowledge base.
func LoadMarketingBible() string {
	// For production, we'd bundle this or read from a consistent path.
	// Here we read from the internal skills directory.
	path := "internal/agent/skills/MARKETING_BIBLE.md"
	data, err := os.ReadFile(path)
	if err != nil {
		return "Marketing Knowledge Base not found."
	}
	return string(data)
}
