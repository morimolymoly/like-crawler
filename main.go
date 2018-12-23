package main

import (
	"fmt"
	"os"

	"github.com/morimolymoly/like-crawler/cmd"
	"github.com/morimolymoly/like-crawler/config"
)

func main() {
	// init config
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd.Execute()
}
