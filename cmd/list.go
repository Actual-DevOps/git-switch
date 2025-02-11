/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Actual-DevOps/git-switch/internal/config"
	"github.com/Actual-DevOps/git-switch/internal/gitprofile"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List git profiles",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		if err := config.LoadConfig(&cfg); err != nil {
			fmt.Printf("Error load config file: %s", err)
		}

		for i := range cfg.Profiles {
			ok, err := gitprofile.IsCurrentProfile(cfg.Profiles[i].Git.User.Email)
			if err != nil {
				fmt.Printf("Error check current profile: %v", err)
			}

			if ok {
				fmt.Printf("\033[32m%d. %s - %s (current)\033[0m\n",
					i+1, cfg.Profiles[i].Name, cfg.Profiles[i].Description)
			} else {
				fmt.Printf("%d. %s - %s\n", i+1, cfg.Profiles[i].Name, cfg.Profiles[i].Description)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
