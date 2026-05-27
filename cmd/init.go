package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yellowhama/musu-marketer/internal/db"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize musu-marketer environment for a project",
	Run: func(cmd *cobra.Command, args []string) {
		project := viper.GetString("project")
		fmt.Printf("🚀 Initializing musu-marketer for project '%s' (Version %s)...\n", project, Version)

		// 1. Create project directories
		baseDir := filepath.Join("projects", project)
		dirs := []string{
			filepath.Join(baseDir, "campaigns"),
			filepath.Join(baseDir, "personas"),
			filepath.Join(baseDir, "data"),
			filepath.Join(baseDir, "published"),
		}
		
		for _, d := range dirs {
			if err := os.MkdirAll(d, 0755); err != nil {
				fmt.Printf("❌ Failed to create directory %s: %v\n", d, err)
				return
			}
			fmt.Printf("✅ Directory ready: ./%s\n", d)
		}

		// 2. Initialize Database for this project
		dbPath := filepath.Join(baseDir, "data", "marketer.db")
		_, err := db.NewStore(dbPath)
		if err != nil {
			fmt.Printf("❌ Failed to initialize database: %v\n", err)
			return
		}
		fmt.Printf("✅ Database ready: %s\n", dbPath)

		// 3. Create default persona if missing
		defaultPersonaPath := filepath.Join(baseDir, "personas", "default.md")
		if _, err := os.Stat(defaultPersonaPath); os.IsNotExist(err) {
			defaultPersona := `Role: Professional AI Marketing Expert
Tone: Informative, authoritative, and slightly enthusiastic.
Framework: AIDA
Formatting: Clean Markdown with bold headers.`
			os.WriteFile(defaultPersonaPath, []byte(defaultPersona), 0644)
			fmt.Printf("✅ Default persona created: %s\n", defaultPersonaPath)
		}

		// 4. Create local project config
		configPath := filepath.Join(baseDir, "config.yaml")
		viper.Set("db_path", dbPath)
		viper.WriteConfigAs(configPath)
		fmt.Printf("✅ Project configuration saved: %s\n", configPath)

		fmt.Printf("\n✨ Project '%s' initialized! Run 'musu-marketer draft [topic] -p %s' to start.\n", project, project)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
