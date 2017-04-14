package joke

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetJoke(t *testing.T) {
	resp := `{"type":"success","value":{"id":0,"joke":"the best joke"}}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(resp))
	}))

	joke := GetJoke(ts.URL)
	t.Log("joke", joke)

	if joke != "the best joke" {
		t.Error("joke is not the best")
	}

	resp = "bad joke "
	joke = GetJoke(ts.URL)
	if joke != "joke error" {
		t.Error("joke must be bad")
	}
}
