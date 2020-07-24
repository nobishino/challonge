package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/nobishino/challonge/challongecli/cmd"
)

func main() {
	cmd.Execute()
}

func playground() {
	apikey := os.Getenv("CHALLONGE_APIKEY")
	baseurl := "https://api.challonge.com/v1/tournaments.json"
	url := fmt.Sprintf("%s?api_key=%s", baseurl, apikey)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
