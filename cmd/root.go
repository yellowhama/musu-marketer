package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const Version = "v2.0.1"

var rootCmd = &cobra.Command{
	Use:     "musu-marketer",
	Short:   "AI Influencer & Marketing Automation Engine",
	Long:    `Transforms verified knowledge from musu-crawl-ai into viral social media content.`,
	Version: Version,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() != "update" && cmd.Name() != "help" {
			go checkNewVersion()
		}
	},
}

func checkNewVersion() {
	latest, _, err := GetLatestRelease("yellowhama", "musu-marketer")
	if err == nil && latest != Version {
		fmt.Fprintf(os.Stderr, "\n💡 New version available: %s (Current: %s)\n", latest, Version)
		fmt.Fprintf(os.Stderr, "👉 Run 'musu-marketer update' to upgrade.\n\n")
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().String("wiki", "C:\\Users\\empty\\musu-crawl-ai\\wiki", "Path to musu-crawl-ai wiki")
	rootCmd.PersistentFlags().StringP("project", "p", "default", "Project name to scope the assets")
	
	viper.BindPFlag("wiki_dir", rootCmd.PersistentFlags().Lookup("wiki"))
	viper.BindPFlag("project", rootCmd.PersistentFlags().Lookup("project"))
}

func initConfig() {
	viper.AutomaticEnv()
}
