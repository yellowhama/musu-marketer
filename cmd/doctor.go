package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yellowhama/musu-marketer/internal/bridge"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check wiki, project, and AI connectivity before drafting campaigns",
	RunE: func(cmd *cobra.Command, args []string) error {
		project := viper.GetString("project")
		wikiDir := viper.GetString("wiki_dir")
		aiProvider := viper.GetString("ai_provider")
		aiURL := viper.GetString("ai_url")
		projectDir := filepath.Join("projects", project)
		dbPath := filepath.Join(projectDir, "data", "marketer.db")
		topic, _ := cmd.Flags().GetString("topic")

		jsonMode := viper.GetBool("json")
		if !jsonMode {
			fmt.Println("==> musu-marketer doctor")
			fmt.Printf("Project      : %s\n", project)
			fmt.Printf("Wiki         : %s\n", wikiDir)
			fmt.Printf("AI Provider  : %s\n", aiProvider)
			fmt.Printf("AI URL       : %s\n", aiURL)
			fmt.Printf("Project Dir  : %s\n", projectDir)
		}

		hasError := false
		report := map[string]interface{}{
			"project":      project,
			"wiki_dir":     wikiDir,
			"ai_provider":  aiProvider,
			"ai_url":       aiURL,
			"project_dir":  projectDir,
			"wiki_exists":  false,
			"wiki_markdown_count": 0,
			"project_exists": false,
			"db_exists":    false,
			"ai_reachable": false,
		}
		if topic != "" {
			report["topic"] = topic
			report["topic_ready"] = false
			report["topic_source_count"] = 0
		}
		autoFix, _ := cmd.Flags().GetBool("fix")

		if info, err := os.Stat(wikiDir); err != nil || !info.IsDir() {
			if !jsonMode {
				fmt.Printf("❌ Wiki directory not found. Expected a musu-crawl-ai wiki at %s\n", wikiDir)
				fmt.Println("   Fix: pass --wiki explicitly or set MUSU_MARKETER_WIKI / MUSU_WIKI.")
			}
			hasError = true
		} else {
			mdCount := countMarkdownFiles(wikiDir)
			if !jsonMode {
				fmt.Printf("✅ Wiki directory exists (%d markdown files)\n", mdCount)
			}
			report["wiki_exists"] = true
			report["wiki_markdown_count"] = mdCount
		}

		if _, err := os.Stat(projectDir); err != nil {
			if autoFix {
				if !jsonMode {
					fmt.Printf("🛠️  Auto-fixing missing project scaffold for %s\n", project)
				}
				if _, _, fixErr := bootstrapProject(project, !jsonMode); fixErr == nil {
					report["project_exists"] = true
					report["db_exists"] = true
				} else {
					report["project_fix_error"] = fixErr.Error()
					if !jsonMode {
						fmt.Printf("❌ Auto-fix failed: %v\n", fixErr)
					}
				}
			} else if !jsonMode {
				fmt.Printf("⚠️  Project directory missing: %s\n", projectDir)
				fmt.Printf("   Run: musu-marketer init --project %s or use doctor --fix\n", project)
			}
		} else {
			if !jsonMode {
				fmt.Println("✅ Project directory exists")
			}
			report["project_exists"] = true
		}

		if _, err := os.Stat(dbPath); err != nil {
			if !jsonMode {
				fmt.Printf("⚠️  Project database missing: %s\n", dbPath)
			}
		} else {
			if !jsonMode {
				fmt.Println("✅ Project database exists")
			}
			report["db_exists"] = true
		}

		if err := probeModels(aiURL); err != nil {
			if !jsonMode {
				fmt.Printf("❌ AI endpoint probe failed: %v\n", err)
			}
			report["ai_error"] = err.Error()
			hasError = true
		} else {
			if !jsonMode {
				fmt.Println("✅ AI endpoint reachable")
			}
			report["ai_reachable"] = true
		}

		if topic != "" && report["wiki_exists"] == true {
			b := bridge.NewWikiBridge(wikiDir)
			sources, topicErr := b.FindByTopic(topic)
			if topicErr != nil {
				report["topic_error"] = topicErr.Error()
				hasError = true
			} else {
				report["topic_source_count"] = len(sources)
				if len(sources) == 0 {
					if !jsonMode {
						fmt.Printf("⚠️  No wiki sources matched topic %q\n", topic)
					}
					hasError = true
				} else {
					if !jsonMode {
						fmt.Printf("✅ Topic %q matched %d wiki source(s)\n", topic, len(sources))
					}
					report["topic_ready"] = true
				}
			}
		}

		if hasError {
			err := fmt.Errorf("doctor found blocking issues")
			printJSONError(err, report)
			return err
		}
		if !jsonMode {
			fmt.Println("✅ Doctor passed")
		}
		printJSONSuccess("Doctor passed", report)
		return nil
	},
}

func init() {
	doctorCmd.Flags().Bool("fix", false, "Auto-create missing local project scaffold when safe")
	doctorCmd.Flags().String("topic", "", "Optional topic to validate draft-readiness against the wiki")
	rootCmd.AddCommand(doctorCmd)
}

func probeModels(baseURL string) error {
	baseURL = strings.TrimRight(strings.TrimSpace(baseURL), "/")
	if baseURL == "" {
		return fmt.Errorf("empty ai-url")
	}
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(baseURL + "/models")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return fmt.Errorf("unexpected status %s from %s/models", resp.Status, baseURL)
}

func countMarkdownFiles(root string) int {
	count := 0
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		if strings.EqualFold(filepath.Ext(path), ".md") {
			count++
		}
		return nil
	})
	return count
}
