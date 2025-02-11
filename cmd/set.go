/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/Actual-DevOps/git-switch/internal/config"
	"github.com/Actual-DevOps/git-switch/internal/gitprofile"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set git profile",
	Run: func(cmd *cobra.Command, args []string) {
		var cfg config.Config
		if err := config.LoadConfig(&cfg); err != nil {
			fmt.Printf("Error load config file: %s", err)
		}

		ok, err := gitprofile.IsProfileNumber(args[0])
		if err != nil {
			fmt.Printf("Error check current profile: %v", err)
		}

		if len(args) == 0 || len(args) > 1 || !ok {
			fmt.Println("Please specify profile number it must be one argument")
		}

		profile, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error convert profile number: %v", err)
		}

		gitprofile.SetProfile(cfg.Profiles, profile)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
