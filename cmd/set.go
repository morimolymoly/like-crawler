package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/morimolymoly/like-crawler/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set consumerkey and accessToken",
	Short: "Set consumerkey and accessToken",
	Long:  `Set consumerkey and accessToken`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.GetInstance()
		fmt.Printf("ConsumerKey:")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		ck := stdin.Text()

		fmt.Printf("AccessToken:")
		stdin.Scan()
		at := stdin.Text()

		err := c.SetConfigs(ck, at)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("ConsumerKey set to %s\nAccessToken set to %s\n", ck, at)
	},
}
