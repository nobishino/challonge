package challonge_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nobishino/challonge"
)

func TestMain(m *testing.M) {
}

func TestIndex(t *testing.T) {
	expect := "hello"
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// GET paramをチェックしたい
		fmt.Fprintln(w, expect)
	})
	ts := httptest.NewServer(h)
	defer ts.Close()
	defer challonge.SetBaseUrl(ts.URL)() // replace API Base URL with test server's base URL

	resp, err := challonge.Get().Index().Do()
	if err != nil {
		t.Error(err)
	}
	if resp != expect+"\n" {
		t.Errorf("want hello but got %s", resp)
	}
}
