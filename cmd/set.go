package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Actual-DevOps/git-switch/internal/config"
	"github.com/Actual-DevOps/git-switch/internal/git"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set git profile",
	Example: "git-switch set 1",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := config.LoadAndValidateConfig(); err != nil {
			fmt.Printf("Config error: %v", err)
			os.Exit(1)
		}

		if len(args) == 0 {
			if err := cmd.Help(); err != nil {
				fmt.Printf("Error print help: %v", err)
				os.Exit(1)
			}
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.ReadConfig()
		if err != nil {
			fmt.Printf("Error read config file: %v", err)
			os.Exit(1)
		}

		profiles, err := git.GetProfiles(conf)
		if err != nil {
			fmt.Println("Error get profile: %w", err)
			os.Exit(1)
		}

		profileNumber, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Enter only profile number!")
			os.Exit(1)
		}

		if profileNumber > len(profiles) {
			fmt.Printf("Profile %d not found\n", profileNumber)
			os.Exit(1)
		}

		confProfilesGit, ok := profiles[profileNumber-1].(map[string]any)["git"].(map[string]any)
		if !ok {
			fmt.Println("Unknow type for profiles.git")
			os.Exit(1)
		}

		flagGlobal, err := cmd.Flags().GetBool("global")
		if err != nil {
			fmt.Println("Error parse 'global' flag: %w", err)
			os.Exit(1)
		}

		if err := git.SetProfile(confProfilesGit, flagGlobal); err != nil {
			fmt.Println("Error set profile: %w", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().BoolP("global", "g", false, "Set Global config")
}
