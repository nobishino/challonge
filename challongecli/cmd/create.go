package cmd

import (
	"log"

	"github.com/nobishino/challonge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create tournament",
	Run:   create,
}

var (
	name      string
	subdomain string
	suffix    string
)

func init() {
	createCmd.Flags().StringVarP(&name, "name", "n", "", "specify tournament name")
	createCmd.Flags().StringVar(&subdomain, "subdomain", "", "specify subdomain")
	createCmd.Flags().StringVar(&suffix, "suffix", "", "specify tournament url suffix")
}

func create(cmd *cobra.Command, args []string) {
	// apikey := os.Getenv(challongeApikeyEnv)
	apikey := viper.GetString("apikey")
	if apikey == "" {
		log.Fatalf("Environment variable %s not set", challongeApikeyEnv)
	}
	s := challonge.NewService(apikey)
	resp, err := s.Create(name, subdomain, suffix)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
}
