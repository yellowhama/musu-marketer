package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yellowhama/musu-marketer/internal/db"
)

func bootstrapProject(project string, verbose bool) (string, string, error) {
	baseDir := filepath.Join("projects", project)
	dirs := []string{
		filepath.Join(baseDir, "campaigns"),
		filepath.Join(baseDir, "personas"),
		filepath.Join(baseDir, "data"),
		filepath.Join(baseDir, "published"),
	}

	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return "", "", fmt.Errorf("create directory %s: %w", d, err)
		}
		if verbose {
			fmt.Printf("✅ Directory ready: ./%s\n", d)
		}
	}

	dbPath := filepath.Join(baseDir, "data", "marketer.db")
	if _, err := db.NewStore(dbPath); err != nil {
		return "", "", fmt.Errorf("initialize database: %w", err)
	}
	if verbose {
		fmt.Printf("✅ Database ready: %s\n", dbPath)
	}

	defaultPersonaPath := filepath.Join(baseDir, "personas", "default.md")
	if _, err := os.Stat(defaultPersonaPath); os.IsNotExist(err) {
		defaultPersona := `Role: Professional AI Marketing Expert
Tone: Informative, authoritative, and slightly enthusiastic.
Framework: AIDA
Formatting: Clean Markdown with bold headers.`
		if err := os.WriteFile(defaultPersonaPath, []byte(defaultPersona), 0644); err != nil {
			return "", "", fmt.Errorf("write default persona: %w", err)
		}
		if verbose {
			fmt.Printf("✅ Default persona created: %s\n", defaultPersonaPath)
		}
	}

	configPath := filepath.Join(baseDir, "config.yaml")
	config := fmt.Sprintf("db_path: %s\nwiki_dir: %s\nai_provider: %s\nai_url: %s\n",
		dbPath,
		viper.GetString("wiki_dir"),
		viper.GetString("ai_provider"),
		viper.GetString("ai_url"),
	)
	if err := os.WriteFile(configPath, []byte(config), 0644); err != nil {
		return "", "", fmt.Errorf("write config: %w", err)
	}
	if verbose {
		fmt.Printf("✅ Project configuration saved: %s\n", configPath)
	}
	return baseDir, dbPath, nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize musu-marketer environment for a project",
	Run: func(cmd *cobra.Command, args []string) {
		project := viper.GetString("project")
		fmt.Printf("🚀 Initializing musu-marketer for project '%s' (Version %s)...\n", project, Version)

		if _, _, err := bootstrapProject(project, true); err != nil {
			fmt.Printf("❌ Failed to initialize project: %v\n", err)
			return
		}

		fmt.Printf("\n✨ Project '%s' initialized! Run 'musu-marketer draft [topic] -p %s' to start.\n", project, project)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
