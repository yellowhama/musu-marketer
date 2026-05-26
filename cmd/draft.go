package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yellowhama/musu-marketer/internal/agent"
	"github.com/yellowhama/musu-marketer/internal/bridge"
	"github.com/yellowhama/musu-marketer/internal/db"
)

var draftCmd = &cobra.Command{
	Use:   "draft [topic]",
	Short: "Draft a viral campaign based on verified knowledge and strategy",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		topic := args[0]
		project := viper.GetString("project")
		wikiDir := viper.GetString("wiki_dir")
		
		dbPath := filepath.Join("projects", project, "data", "marketer.db")
		model, _ := cmd.Flags().GetString("model")
		persona, _ := cmd.Flags().GetString("persona")

		ExecuteDraft(topic, wikiDir, dbPath, model, persona, project)
	},
}

func ExecuteDraft(topic, wikiDir, dbPath, model, persona, project string) error {
	fmt.Printf("🌉 Connecting to Wiki at: %s (Project Scope: %s)\n", wikiDir, project)
	b := bridge.NewWikiBridge(wikiDir)
	
	sources, err := b.FindByTopic(topic)
	if err != nil || len(sources) == 0 {
		return fmt.Errorf("no verified knowledge found for topic: %s", topic)
	}

	fmt.Printf("📚 Found %d knowledge sources. Starting Strategic Brain...\n", len(sources))
	
	var context strings.Builder
	for _, s := range sources {
		context.WriteString(fmt.Sprintf("--- Source: %s ---\n%s\n\n", s.Title, s.Content))
	}

	// 1. Generate Strategy Brief
	fmt.Println("🧠 Phase 1: Analyzing Market Strategy (STP & Triggers)...")
	strategist := agent.NewStrategist("", model)
	brief, err := strategist.CreateBrief(context.String())
	if err != nil {
		fmt.Printf("   ⚠️  Strategy phase failed: %v. Using default approach.\n", err)
		brief = &agent.MarketingBrief{ValueProp: "General information", Target: "General audience"}
	} else {
		fmt.Printf("   🎯 [Brief] Target: %s | Goal: %s\n", brief.Target, brief.Goal)
		fmt.Printf("   🔥 Triggers: %s\n", strings.Join(brief.Triggers, ", "))
	}

	// 2. Generate Campaign Content
	fmt.Printf("✍️  Phase 2: Drafting content using Persona: %s\n", persona)
	projectPath := filepath.Join("projects", project)
	copywriter := agent.NewCopywriter("", model, persona, projectPath)
	campaign, err := copywriter.GenerateCampaign(topic, context.String(), brief)
	if err != nil {
		return fmt.Errorf("failed to generate campaign: %v", err)
	}

	// 3. Save to Database
	store, err := db.NewStore(dbPath)
	if err == nil {
		briefJSON, _ := json.Marshal(brief)
		id, _ := store.SaveCampaign(topic, campaign, string(briefJSON), persona)
		fmt.Printf("\n✅ Strategic Campaign saved to database (ID: %d)\n", id)
	}

	fmt.Println("\n🏁 --- GENERATED STRATEGIC CAMPAIGN --- 🏁")
	fmt.Println(campaign)
	return nil
}

func init() {
	draftCmd.Flags().String("model", "llama3", "Ollama model for reasoning")
	draftCmd.Flags().String("persona", "default", "Persona to use for this draft")
	rootCmd.AddCommand(draftCmd)
}
