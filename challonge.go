package challonge

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var baseUrl = "https://api.challonge.com/v1"

type Interface interface {
	Create(title string)
}

type Service struct {
	apikey  string
	baseUrl string
}
type CreateParam struct {
	ApiKey     string     `json:"api_key"`
	Tournament Tournament `json:"tournament"`
}

type Tournament struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Subdomain string `json:"subdomain"`
}

func (s Service) Create(name, subdomain, tournamentPath string) (string, error) {
	path := "tournaments.json"
	url := fmt.Sprintf("%s/%s", s.baseUrl, path)
	if tournamentPath == "" {
		now := time.Now().Format(time.RFC3339)
		hash := sha256.Sum256([]byte(subdomain + now))
		tournamentPath = string(hash[:])
	}
	p := CreateParam{
		ApiKey: s.apikey,
		Tournament: Tournament{
			Name:      name,
			Url:       tournamentPath,
			Subdomain: subdomain,
		},
	}
	b, err := json.MarshalIndent(p, "", "  ")
	log.Println("POSTING:")
	log.Println(string(b))
	if err != nil {
		return "", err
	}
	body := bytes.NewReader(b)
	resp, err := http.Post(url, "application/json", body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBytes), nil
}

func NewService(apikey string) Service {
	return Service{
		apikey:  apikey,
		baseUrl: baseUrl,
	}
}
