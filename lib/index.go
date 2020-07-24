package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GetRequest struct {
	apikey string
	url    string
	path   string
}

func Get() GetRequest {
	return GetRequest{
		apikey: os.Getenv("CHALLONGE_APIKEY"),
		url:    "",
	}
}

func (g GetRequest) Index() GetRequest {
	g.path = "tournaments.json"
	return g
}

func (g GetRequest) WithURL(url string) GetRequest {
	g.url = url
	return g
}

func (g GetRequest) Do() (string, error) {
	url := fmt.Sprintf("%s/%s?api_key=%s", g.url, g.path, g.apikey)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP Status Code %v", resp.StatusCode)
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
