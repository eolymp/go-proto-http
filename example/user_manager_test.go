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
		resp, err := http.Get(srv.URL + "/users/sk10y/comments?sort=10&order=ASC")
		if err != nil {
			t.Fatal(err)
		}

		if got, want := resp.StatusCode, http.StatusNotImplemented; got != want {
			t.Errorf("Response status code does not match expected value: got %v, want %v", got, want)
		}
	})
}
