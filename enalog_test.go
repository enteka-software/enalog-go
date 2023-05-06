package enalog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPushEvent(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimSpace(r.URL.Path) {
		case "/v1/events":
			fmt.Fprintf(w, "")
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))
	defer s.Close()

	client := New("123")
	client.BaseUrl = s.URL

	event := Event{
		Project:     "landing-page",
		Name:        "user-joined-waitlist",
		Push:        true,
		Description: "User joined the waitlist",
		Icon:        "🚀",
		Tags:        []string{},
		Meta:        map[string]string{},
		Channels:    map[string]string{},
	}

	res, err := client.PushEvent(event)

	statusCode := res["statusCode"]
	if statusCode != "200" {
		t.Errorf("received a status code was not 200, was: %v instead", err)
	}

	if err != nil {
		t.Errorf("received error: %v", err)
	}
}
