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
		BaseUrl:  "https://api.enalog.app",
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
		log.Println(err)
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

// CheckFeature
func (client *Client) CheckFeature(flag FeatureFlag) bool {
	f, _ := json.Marshal(flag)
	url := client.BaseUrl + "/v1/feature-flags"
	authHeader := "Bearer " + client.apiToken

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(f))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authHeader)

	res, err := client.http.Do(req)
	if err != nil {
		log.Println(err)
		return false
	}

	defer res.Body.Close()
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return false
	}

	body := string(b)

	if res.StatusCode == 200 {
		flagRes := FeatureFlagRes{}
		err = json.Unmarshal([]byte(body), &flagRes)
		if err != nil {
			return false
		}

		if flagRes.FlagType == "Boolean" {
			if flagRes.Variant == "a-variant" {
				return true
			} else {
				return false
			}
		}
	}

	errorMessage := fmt.Sprintf("Error %v - %s", res.StatusCode, body)
	log.Println(errorMessage)
	return false
}
