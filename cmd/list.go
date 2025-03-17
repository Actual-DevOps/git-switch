package cmd

import (
	"fmt"
	"os"

	"github.com/Actual-DevOps/git-switch/internal/config"
	"github.com/Actual-DevOps/git-switch/internal/git"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List git profiles",
	PreRun: func(_ *cobra.Command, _ []string) {
		if err := config.LoadAndValidateConfig(); err != nil {
			fmt.Printf("Config error: %v", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, _ []string) {
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

		for i := range profiles {
			gitProfiles, ok := profiles[i].(map[string]any)
			if !ok {
				fmt.Printf("Error parse gitProfiles")
				os.Exit(1)
			}

			gitProfilesValues, ok := gitProfiles["git"].(map[string]any)
			if !ok {
				fmt.Printf("Error parse gitProfilesValues")
				os.Exit(1)
			}

			flagExtend, err := cmd.Flags().GetBool("extend")
			if err != nil {
				fmt.Println("Error parse 'extend' flag: %w", err)
				os.Exit(1)
			}

			var startColor, resetColor string
			for k, v := range gitProfilesValues {
				if k == "user.email" {
					isCurrentProfile, err := git.IsCurrentProfile(v.(string))
					if err != nil {
						fmt.Printf("Error IsCurrentProfile: %v", err)
						os.Exit(1)
					}

					if isCurrentProfile {
						startColor = git.GreenColor
						resetColor = git.ResetColor
					}
				}
			}

			fmt.Printf(startColor+"%d. %s - %s\n"+resetColor, i+1, gitProfiles["name"], gitProfiles["description"])

			if flagExtend {
				for k, v := range gitProfilesValues {
					fmt.Printf("\t%s = \"%s\"\n", k, v)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("extend", "e", false, "Show all options for profiles")
}
