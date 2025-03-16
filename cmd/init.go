package cmd

import (
	"fmt"
	"os"

	"github.com/Actual-DevOps/git-switch/internal/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("Create simple config in ~/%s", config.ConfigFileName),
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", os.Getenv("HOME"), config.ConfigFileName)); err == nil {
			fmt.Println("Config file already exists!")
			os.Exit(1)
		}

		err := os.WriteFile(
			fmt.Sprintf("%s/%s", os.Getenv("HOME"), config.ConfigFileName),
			[]byte(`
profiles:
  - name: default
    description: Default git profile
    git:
      user.name: John Doe
      user.email: john.doe@example.com
      core.editor: vim
      init.defaultBranch: master`),0644)
		if err != nil {
			fmt.Println("Error create config file: %w", err)
			os.Exit(1)
		}

		fmt.Printf("Config file ~/%s created!\n", config.ConfigFileName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
