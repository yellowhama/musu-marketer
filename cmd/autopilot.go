package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var autopilotCmd = &cobra.Command{
	Use:   "autopilot [topic/subreddit]",
	Short: "Automatically spot trends, research, and draft campaigns",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		project := viper.GetString("project")
		wikiDir := viper.GetString("wiki_dir")
		dbPath := filepath.Join("projects", project, "data", "marketer.db")
		model, _ := cmd.Flags().GetString("model")
		
		// Decoupled: Read crawl path from config, default to relative path
		crawlPath := viper.GetString("crawl_path")
		if crawlPath == "" {
			crawlPath = "../musu-crawl-ai/musu-crawl.exe"
		}

		crawlProject := "autopilot-" + project

		fmt.Printf("🚀 Starting Autopilot for topic '%s' in project '%s'\n", topic, project)
		fmt.Printf("[DEBUG] Using crawl binary: %s\n", crawlPath)

		// 1. Spot Trend and Trigger Research
		fmt.Println("🔍 Step 1: Spotting trends and starting research...")
		spotArgs := []string{"spot", topic, "--limit", "1", "--research", "--project", crawlProject}
		
		spotCmd := exec.Command(crawlPath, spotArgs...)
		stdout, _ := spotCmd.StdoutPipe()
		spotCmd.Start()

		scanner := bufio.NewScanner(stdout)
		var topTrendTitle string
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println("   [CRAWL]", line)
			if strings.Contains(line, "Auto-triggering research for top trend:") {
				re := regexp.MustCompile(`"([^"]+)"`)
				match := re.FindStringSubmatch(line)
				if len(match) > 1 {
					topTrendTitle = match[1]
				}
			}
		}
		spotCmd.Wait()

		if topTrendTitle == "" {
			fmt.Println("❌ Could not identify a top trend to research.")
			return
		}

		// 2. Deep researching
		fmt.Printf("\n🧠 Step 2: Deep researching top trend: %q\n", topTrendTitle)
		researchArgs := []string{"research", topTrendTitle, "--project", crawlProject, "--depth", "2"}
		researchCmd := exec.Command(crawlPath, researchArgs...)
		researchCmd.Stdout = os.Stdout
		researchCmd.Stderr = os.Stderr
		if err := researchCmd.Run(); err != nil {
			fmt.Printf("❌ Research failed: %v\n", err)
			return
		}

		// 3. Draft Campaigns
		fmt.Println("\n✍️  Step 3: Drafting campaigns for personas...")
		personaDir := filepath.Join("projects", project, "personas")
		files, _ := os.ReadDir(personaDir)
		
		var personas []string
		for _, f := range files {
			if filepath.Ext(f.Name()) == ".md" {
				personas = append(personas, strings.TrimSuffix(f.Name(), ".md"))
			}
		}
		
		if len(personas) == 0 {
			personas = append(personas, "default")
		}

		for _, p := range personas {
			fmt.Printf("\n--- Drafting for Persona: %s ---\n", p)
			if err := ExecuteDraft(topTrendTitle, wikiDir, dbPath, model, p, project); err != nil {
				fmt.Printf("   ⚠️  Failed to draft for %s: %v\n", p, err)
			}
		}

		fmt.Printf("\n🏁 Autopilot mission complete for project '%s'.\n", project)
	},
}

func init() {
	autopilotCmd.Flags().String("model", "llama3", "Ollama model for copywriting")
	rootCmd.AddCommand(autopilotCmd)
}
