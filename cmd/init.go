package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create simple config in ~/.git-switch.conf",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("here need to create config")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
