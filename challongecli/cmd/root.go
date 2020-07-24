package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const challongeApikeyEnv = "CHALLONGE_APIKEY"

func init() {
	rootCmd.AddCommand(createCmd)
}

var rootCmd = &cobra.Command{
	Use:   "challonge",
	Short: "access challonge API",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
