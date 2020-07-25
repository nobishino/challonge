package challonge

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// GetRequest represents GET for API
type GetRequest struct {
	apikey string
	url    string
	path   string
}

// Get returns default GetRequest
func Get() GetRequest {
	return GetRequest{
		apikey: os.Getenv("CHALLONGE_APIKEY"),
		url:    baseUrl,
	}
}

// Index returns GetRequest to index endpoint
func (g GetRequest) Index() GetRequest {
	g.path = "tournaments.json"
	return g
}

// WithURL returns GetRequest with custom base url for testing
func (g GetRequest) WithURL(url string) GetRequest {
	g.url = url
	return g
}

// Do executes Request
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
