package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "show liked list",
	Short: "Show liked list",
	Long:  `Show liked list`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
