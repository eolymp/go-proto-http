package compete

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewUserManagerHandler(t *testing.T) {
	srv := httptest.NewServer(NewUserManagerHandler(
		&UnimplementedUserManagerServer{},
	))

	defer srv.Close()

	t.Run("parse query parameters", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "/users/sk10y/comments?offset=10&order=ASC")
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusNotImplemented; got != want {
			t.Errorf("Response status code does not match expected value: got %v, want %v", got, want)
		}
	})

	t.Run("parse invalid query parameters", func(t *testing.T) {
		resp, err := http.Get(srv.URL + "/users/sk10y/comments?offset=xx1&order=ASC")
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusBadRequest; got != want {
			t.Errorf("Response status code does not match expected value: got %v, want %v", got, want)
		}
	})
}
