package main

import (
	"fmt"
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
		fmt.Printf("Error fetching events: %v\n", err)
		return
	}
	displayActivity(events)
}
