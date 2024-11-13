package main

import "fmt"

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
