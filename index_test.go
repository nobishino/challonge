package challonge_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nobishino/challonge"
)

func TestIndex(t *testing.T) {
	expect := "hello"
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// GET paramをチェックしたい
		fmt.Fprintln(w, expect)
	})
	ts := httptest.NewServer(h)
	defer ts.Close()

	resp, err := challonge.Get().WithURL(ts.URL).Index().Do()
	if err != nil {
		t.Error(err)
	}
	if resp != expect+"\n" {
		t.Errorf("want hello but got %s", resp)
	}
}
