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
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.ReadConfig()
		if err != nil {
			fmt.Printf("Error read config file: %v", err)
			os.Exit(1)
		}


		// git.IsValidConfig(conf)

		// os.Exit(1)

		profiles, err := git.GetProfiles(conf)
		if err != nil {
			fmt.Println("Error get profile: %v", err)
			os.Exit(1)
		}

		for i := range profiles {
			gitProfiles, ok := profiles[i].(map[string]any)
			if !ok {
				fmt.Printf("Error parse profiles.git: %v", err)
				os.Exit(1)
			}

			gitProfilesValues, ok := gitProfiles["git"].(map[string]any)
			if !ok {
				fmt.Printf("Error parse profiles.git: %v", err)
				os.Exit(1)
			}

			flagExtend, err := cmd.Flags().GetBool("extend")
			if err != nil {
				fmt.Println("Error parse 'extend' flag: %w", err)
				os.Exit(1)
			}

			fmt.Printf("%d. %s - %s\n", i+1, gitProfiles["name"], gitProfiles["description"])

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
