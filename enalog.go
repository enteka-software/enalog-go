package enalog

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	apiToken string
	http     http.Client
	BaseUrl  string
}

func New(apiToken string) *Client {
	c := &Client{
		apiToken: apiToken,
		http:     makeHttpClient(),
		BaseUrl:  "http://127.0.0.1:3000",
	}

	return c
}

func makeHttpClient() http.Client {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	return httpClient
}

// PushEvent
func (client *Client) PushEvent(event Event) (map[string]string, error) {
	e, _ := json.Marshal(event)
	url := client.BaseUrl + "/v1/events"
	authHeader := "Bearer " + client.apiToken

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(e))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	res, err := client.http.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("error sending HTTP request to EnaLog")
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("error reading request body from EnaLog")
	}

	body := string(b)

	if res.StatusCode == 200 {
		return map[string]string{"statusCode": "200", "message": body}, nil
	}

	errorMessage := fmt.Sprintf("Error %v - %s", res.StatusCode, body)
	return nil, errors.New(errorMessage)
}
