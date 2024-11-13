package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}

	username := os.Args[1]
	fmt.Printf("Fetching recent activity for user: %s...\n", username)
	events, err := githubGetEvent(username)
	if err != nil {
		fmt.Printf("Error fetching events: %s\n", err)
	}
	displayActivity(events)
}

var eventHandlers = map[string]func(event GitHubEvent){
	"PushEvent": func(event GitHubEvent) {
		if len(event.Payload.Commits) > 0 {
			fmt.Printf("Pushed %d commits to %s\n", len(event.Payload.Commits), event.Repo.Name)
		} else {
			fmt.Printf("Pushed to %s\n", event.Repo.Name)
		}
	},
	"WatchEvent": func(event GitHubEvent) {
		fmt.Printf("Starred repository %s\n", event.Repo.Name)
	},
	"IssuesEvent": func(event GitHubEvent) {
		fmt.Printf("Opened an issue in %s\n", event.Repo.Name)
	},
	"ForkEvent": func(event GitHubEvent) {
		fmt.Printf("Forked repository %s\n", event.Repo.Name)
	},
	"CreateEvent": func(event GitHubEvent) {
		fmt.Printf("Created a new repository %s\n", event.Repo.Name)
	},
	"PullRequestEvent": func(event GitHubEvent) {
		fmt.Printf("Opened a pull request in %s\n", event.Repo.Name)
	},
	"DeleteEvent": func(event GitHubEvent) {
		fmt.Printf("Deleted something in %s\n", event.Repo.Name)
	},
}

func displayActivity(events []GitHubEvent) {
	if events == nil {
		fmt.Println("No events found.")
		return
	}

	for _, event := range events {
		if handler, exists := eventHandlers[event.Type]; exists {
			handler(event)
		} else {
			fmt.Printf("Performed %s on %s\n", event.Type, event.Repo.Name)
		}
	}
}

const apiURL = "https://api.github.com/users/%s/events"

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
