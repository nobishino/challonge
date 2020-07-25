package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(indexCmd)
}

var indexCmd = &cobra.Command{
	Use:   "index",
	Short: "retrieve all tournaments",
	Run:   index,
}

func index(c *cobra.Command, args []string) {
	apikey := os.Getenv("CHALLONGE_APIKEY")
	baseurl := "https://api.challonge.com/v1/tournaments.json"
	url := fmt.Sprintf("%s?api_key=%s", baseurl, apikey)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("Got response:")
	fmt.Println(string(body))
}
