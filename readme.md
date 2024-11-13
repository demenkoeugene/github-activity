# GitHub Activity CLI Tool

## 📄 Project Overview

This project is a simple **Command Line Interface (CLI)** tool that allows you to fetch and display the **recent activity of a GitHub user** directly in the terminal. It demonstrates how to work with the GitHub API, handle JSON data, and build a CLI application using the Go programming language.

## 🛠️ Features

- Fetches the recent activity of a specified GitHub user.
- Displays different types of activities, such as:
    - Pushed commits to a repository.
    - Opened issues or pull requests.
    - Starred repositories.
    - Forked repositories.
- Handles errors gracefully, such as invalid usernames or API failures.
- Easy-to-use CLI interface.

## 🖥️ Usage

### Prerequisites
- Make sure you have **Go** installed on your system. If not, download it from [Go's official site](https://golang.org/dl/).

### Installation

1. Clone this repository to your local machine:
    ```bash
    git clone https://github.com/your-username/github-activity.git
    cd github-activity
    ```

2. Build the project:
    ```bash
    go build -o github-activity
    ```

### Running the CLI Tool

To fetch the recent activity of a GitHub user, run:

```bash
./github-activity <username>
```

For example:

```bash
./github-activity demenkoeugene
```

Sample Output
```bash
Fetching recent activity for user: demenkoeugene...
Starred repository avelino/awesome-go
Performed PublicEvent on demenkoeugene/task-tracker
Starred repository go-ozzo/ozzo-validation
Starred repository mtdvio/every-programmer-should-know
```

### 📁 Project Structure
```bash
github-activity/
├── main.go
├── activity.go
├── events.go
├── go.mod
└── README.md
```
* `main.go`: Entry point for the application.
* `activity.go`: Handles fetching data from the GitHub API.
* `events.go`: Handles different types of events and displays them.
* `mod`: Go module file.

### 📝 How It Works

1. Fetches Data: The tool uses the GitHub API to fetch recent activities using the endpoint:
```bash
https://api.github.com/users/<username>/events
```
2. Displays Activity: It processes the JSON response and displays events like pushes, issues, pull requests, and starred repositories. 
3. Handles Errors: If the username is invalid or there are issues with the network, the tool displays appropriate error messages.

📚 Dependencies

* Standard Go Libraries:
  * net/http for making HTTP requests.
  * encoding/json for parsing JSON responses.	
  * fmt and os for input/output operations.

This project does not use any external libraries.

### ⚠️ Limitations
* The GitHub API has rate limits for unauthenticated requests:	
* 60 requests per hour for unauthenticated users. 
* You can increase this limit by using a GitHub Personal Access Token.

### 📄 License

This project is licensed under the MIT License. Feel free to use, modify, and distribute it as you wish.

### Author

Developed by Yevhenii Demenko.
Feel free to reach out on [LinkedIn](https://linkedin.com/in/demenkoeugene)
