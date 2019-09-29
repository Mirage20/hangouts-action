package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/mirage20/hangouts-action/github"
	"github.com/mirage20/hangouts-action/hangouts"
)

var (
	webhookUrl = "unknown"
)

func main() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
	githubToken := getEnvOrFail("GITHUB_TOKEN")
	githubRepo := getEnvOrFail("GITHUB_REPOSITORY")
	githubEventPath := getEnvOrFail("GITHUB_EVENT_PATH")
	// Incorrect sha for forced push
	// githubSha := getEnvOrFail("GITHUB_SHA")

	event := loadEvent(githubEventPath)

	ghc := github.NewClient(githubToken, githubRepo)
	hc := hangouts.NewWebhookClient(webhookUrl)

	_, err := ghc.Validate()
	if err != nil {
		log.Fatal(err)
	}
	githubSha := event.PullRequest.Head.Sha
	var title string
	switch event.Action {
	case "opened":
		title = "New pull request is opened"
	case "reopened":
		title = "Pull request re-opened"
	case "synchronize":
		title = "Pull request updated"
	default:
		return
	}
	sendHangoutMessage(hc, fmt.Sprintf("%d", event.PullRequest.Number), makeMessageFromPullRequest(&event.PullRequest, title, getGitHubStatusResponse(ghc, githubSha), getGitHubCheckRunsResponse(ghc, githubSha)))

	for {
		time.Sleep(5 * time.Second)

		statusResp := getGitHubStatusResponse(ghc, githubSha)
		statusMap := make(map[Status]bool)
		for _, v := range statusResp.Statuses {
			status := statusFromGithubStatus(&v)
			statusMap[status] = true
		}

		checkResp := getGitHubCheckRunsResponse(ghc, githubSha)

		for _, v := range checkResp.CheckRuns {
			// Skip our own check
			if v.Name == "Hangouts" {
				continue
			}
			status := statusFromGithubCheckRun(&v)
			statusMap[status] = true
		}
		overallStatus := getOverallStatus(statusMap)
		if overallStatus == StatusFailure || overallStatus == StatusSuccess {
			checkTitle := "All checks have passed"
			if overallStatus == StatusFailure {
				checkTitle = "Some checks were not successful"
			}
			sendHangoutMessage(hc, fmt.Sprintf("%d", event.PullRequest.Number), makeMessageFromPullRequest(&event.PullRequest, checkTitle, statusResp, checkResp))
			break
		}
	}
}

func getEnvOrFail(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("environment variable %s not provided", key)
	}
	return v
}

func getGitHubStatusResponse(ghc *github.GitHub, sha string) *github.StatusResponse {
	resp, err := ghc.Status(sha)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", resp)
	return resp
}

func getGitHubCheckRunsResponse(ghc *github.GitHub, sha string) *github.CheckRunsResponse {
	resp, err := ghc.CheckRuns(sha)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", resp)
	return resp
}

func sendHangoutMessage(hc *hangouts.Hangouts, thread string, msg *hangouts.Message) {
	m, err := hc.Send(thread, msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", m)
	log.Printf("%+v\n", m.Thread)
}

func loadEvent(eventPath string) *github.Event {
	eventData, err := ioutil.ReadFile(eventPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(eventData))
	event := &github.Event{}
	err = json.Unmarshal(eventData, event)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", event)
	return event
}
