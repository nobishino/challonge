package cmd

import (
	"fmt"
	"os"

	"github.com/nobishino/challonge/lib/index"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "challonge",
	Short: "challonge is a cli client for challonge API",
	// Long: `A Fast and Flexible Static Site Generator built with
	// 			  love by spf13 and friends in Go.
	// 			  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("challonge")
		s := index.NewService()
		resp := s.Get()
		fmt.Println(resp)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
