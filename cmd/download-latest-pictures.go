package cmd

import (
	"fmt"
	"os"

	"github.com/morimolymoly/like-crawler/client"
	"github.com/morimolymoly/like-crawler/config"
	"github.com/morimolymoly/like-crawler/downloader"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dlpCmd)
}

var dlpCmd = &cobra.Command{
	Use:   "download-latest-pictures [screenName]",
	Short: "Download all pictures which included in latestLiked tweets",
	Long:  `Download all pictures which included in latest Liked tweets`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("tell your screen name!")
			return
		}

		cf := config.GetInstance()
		err := client.Init()
		if err != nil {
			fmt.Println(err)
			return
		}

		c := client.GetInstance()
		sn := args[0]
		lurl, err := c.GetLatestLikedPictureURLs(sn)

		if err != nil {
			fmt.Println(err)
			return
		}

		errs := downloader.DownloadAll(lurl)
		if len(errs) == 0 {
			return
		}
		for _, e := range errs {
			fmt.Println(e)
		}
		os.Exit(1)
	},
}
