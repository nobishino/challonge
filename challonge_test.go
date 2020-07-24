package challonge_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/nobishino/challonge"
)
func TestCreate(t *testing.T) {
	expect := "create"
	var reqBody challonge.CreateParam
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expect)
		defer r.Body.Close()
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}
		if err := json.Unmarshal(b, &reqBody); err != nil {
			return
		}
	})
	ts := httptest.NewServer(h)
	defer ts.Close()
	defer challonge.SetBaseUrl(ts.URL)()

	apikey:="dummy_apikey"

	s := challonge.NewService(apikey)
	name := "Test Tournament"
	resp, err := s.Create(name)
	if err != nil {
		t.Error(err)
	}
	if resp != expect {
		t.Errorf("want %s, got %s", expect, resp)
	}
	if reqBody.Tournament.Name != name {
		t.Errorf("want %s, got %s", name, reqBody.Tournament.Name)
	}
	if reqBody.ApiKey != apikey {
		t.Errorf("want %s, got %s", apikey, reqBody.ApiKey)
	}
}

func TestCreateParam(t *testing.T) {
	apikey := "dummy-apikey"
	name := "dummy-tournament-name"
	p := challonge.CreateParam{
		ApiKey: apikey,
		Tournament: challonge.Tournament{
			Name: name,
		},
	}
	jsonBytes, err := json.MarshalIndent(p,"","  ")
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(string(jsonBytes),"api_key") {
		t.Error(string(jsonBytes))
	}
}