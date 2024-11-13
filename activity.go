package main

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
