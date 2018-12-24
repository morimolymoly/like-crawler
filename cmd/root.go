package cmd

import (
	"fmt"
	"os"

	"github.com/morimolymoly/like-crawler/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "like-crawler",
	Short: "Crawler for Twitter's Liked awesome pictures",
	Long:  `Crawler for Twitter's Liked awesome pictures`,
	Run: func(cmd *cobra.Command, args []string) {
		// check config states
		c := config.GetInstance()
		c.ReadConfig()
		err := config.CheckConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute ... execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
