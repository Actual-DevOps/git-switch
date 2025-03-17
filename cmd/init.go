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
	Run: func(_ *cobra.Command, _ []string) {
		if _, err := os.Stat(fmt.Sprintf("%s/%s", os.Getenv("HOME"), config.ConfigFileName)); err == nil {
			fmt.Println("Config file already exists!")
			os.Exit(1)
		}

		err := os.WriteFile(
			fmt.Sprintf("%s/%s", os.Getenv("HOME"), config.ConfigFileName),
			[]byte(config.ExampleConfig),
			0600,
		)

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
