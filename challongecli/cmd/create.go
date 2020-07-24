package cmd

import (
	"log"
	"os"

	"github.com/nobishino/challonge"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create tournament",
	Run:   create,
}

func create(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalln("should enter tournament name")
		os.Exit(1)
	}
	name := args[0]
	apikey := os.Getenv(challongeApikeyEnv)
	if apikey == "" {
		log.Fatalf("Environment variable %s not set", challongeApikeyEnv)
	}
	s := challonge.NewService(apikey)
	resp, err := s.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}
