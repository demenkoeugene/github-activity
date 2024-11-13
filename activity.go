package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.github.com/users/%s/events"

type GitHubEvent struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	CreatedAt string  `json:"created_at"`
}

type Actor struct {
	ID          int    `json:"id"`
	Login       string `json:"login"`
	DisplayName string `json:"display_login"`
	AvatarURL   string `json:"avatar_url"`
	URL         string `json:"url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Payload struct {
	Action  string `json:"action"`
	Commits []struct {
		Message string `json:"message"`
	} `json:"commits"`
}

func githubGetEvent(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf(apiURL, username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	var githubEvents []GitHubEvent
	err = json.NewDecoder(resp.Body).Decode(&githubEvents)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return githubEvents, nil
}
