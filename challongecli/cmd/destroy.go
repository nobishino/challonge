package cmd

import (
	"log"

	"github.com/nobishino/challonge"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy tournament with given url",
	Run:   destroy,
}

func init() {
	destroyCmd.Flags().StringVar(&subdomain, "subdomain", "", "specify subdomain")
	destroyCmd.Flags().StringVar(&suffix, "suffix", "", "specify tournament url suffix")
}

func destroy(cmd *cobra.Command, args []string) {
	apikey := getApiKey()
	if apikey == "" {
		log.Fatalln("api key is not given")
	}
	s := challonge.NewService(apikey)
	s.Destroy(subdomain, suffix)
}
