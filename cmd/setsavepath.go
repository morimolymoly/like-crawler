package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/morimolymoly/like-crawler/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(spCmd)
}

var spCmd = &cobra.Command{
	Use:   "setsavepath [path]",
	Short: "Set savepath",
	Long:  `Set savepath`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.GetInstance()
		fmt.Printf("SavePath:")
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		sp := stdin.Text()

		err := c.UpdateSavePath(sp)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("savePath set to %s\n", sp)
	},
}
