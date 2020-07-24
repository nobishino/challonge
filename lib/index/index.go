package index

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Service は、APIへのアクセスを提供する
type Service interface {
	Get(
		apikey string,
	) string
}

type ServiceImpl struct {
	apikey string
}

func (s ServiceImpl) Get() string {
	baseurl := "https://api.challonge.com/v1/tournaments.json"
	url := fmt.Sprintf("%s?api_key=%s", baseurl, s.apikey)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}

func NewService() ServiceImpl {
	apikey := os.Getenv("CHALLONGE_APIKEY")
	return ServiceImpl{
		apikey: apikey,
	}
}
