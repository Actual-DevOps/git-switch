/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Actual-DevOps/git-switch/internal/config"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List git profiles",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		if err := config.LoadConfig(&cfg); err != nil {
			fmt.Printf("Error load config file: %s", err)
		}

		for i := range cfg.Profiles {
			fmt.Printf("%d. %s - %s\n", i+1, cfg.Profiles[i].Name, cfg.Profiles[i].Description)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
