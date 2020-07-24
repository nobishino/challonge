package challonge

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (s Service) Create(name string) (string, error) {
	path := "tournaments.json"
	url := fmt.Sprintf("%s/%s", s.baseUrl, path)
	p := CreateParam{
		ApiKey: s.apikey,
		Tournament: Tournament{
			Name: name,
			Url:  "asodfjoasodf",
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
