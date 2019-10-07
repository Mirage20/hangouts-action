package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/go-github/v28/github"
	// "github.com/mirage20/hangouts-action/github"
	"github.com/mirage20/hangouts-action/hangouts"
	"golang.org/x/oauth2"
)

var (
	webhookUrl = "unknown"
)

func main() {
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}
	githubToken := getEnvOrFail("GITHUB_TOKEN")
	// githubRepo := getEnvOrFail("GITHUB_REPOSITORY")
	githubEventPath := getEnvOrFail("GITHUB_EVENT_PATH")
	// Incorrect sha for forced push
	// githubSha := getEnvOrFail("GITHUB_SHA")
	event := loadEvent(githubEventPath)

	if skipLabel, ok := os.LookupEnv("SKIP_NOTIFY_LABEL"); ok {
		for _, l := range event.PullRequest.Labels {
			if *l.Name == skipLabel {
				return
			}
		}
	}

	ctx := context.Background()
	ghc := github.NewClient(oauth2.NewClient(
		ctx,
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken}),
	))
	hc := hangouts.NewWebhookClient(webhookUrl)
	ha := &HangoutsAction{
		githubClient:   ghc,
		hangoutsClient: hc,
		SelfActionName: getEnvOrFail("SELF_ACTION_NAME"),
	}

	err := ha.NotifyPullRequest(event, func(event *github.PullRequestEvent) bool {
		return true
	})
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(15 * time.Second)
		done := false
		err = ha.NotifyPullRequestChecks(event, func(event *github.PullRequestEvent, checks Checks) bool {
			if checks.OverallStatus() == StatusFailure || checks.OverallStatus() == StatusSuccess {
				done = true
				return true
			}
			// checks are in progress. no need to send message
			return false
		})
		if err != nil {
			log.Fatal(err)
		}
		if done {
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

func loadEvent(eventPath string) *github.PullRequestEvent {
	eventData, err := ioutil.ReadFile(eventPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(eventData))
	event := &github.PullRequestEvent{}
	err = json.Unmarshal(eventData, event)
	if err != nil {
		log.Fatal(err)
	}
	return event
}
