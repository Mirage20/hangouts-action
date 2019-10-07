package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v28/github"
	"github.com/mirage20/hangouts-action/hangouts"
)

const (
	ImageSuccess      = "https://www.shareicon.net/download/2017/02/09/878601_check_512x512.png"
	ImageFailure      = "https://www.shareicon.net/download/2017/02/09/878603_close_512x512.png"
	ImageInProgress   = "https://www.shareicon.net/download/2017/02/09/878594_gear_512x512.png"
	ImageGitHubAvatar = "https://avatars0.githubusercontent.com/in/15368?s=40&v=4"
)
const (
	StatusSuccess    Status = "Success"
	StatusFailure    Status = "Failure"
	StatusInProgress Status = "InProgress"
)

type Status string

type Check struct {
	Name      string
	Message   string
	TargetUrl string
	AvatarUrl string
	Status    Status
}

type Checks map[Status][]Check

type PullRequestFilter func(event *github.PullRequestEvent) bool
type PullRequestChecksFilter func(event *github.PullRequestEvent, checks Checks) bool
type HangoutsAction struct {
	githubClient   *github.Client
	hangoutsClient *hangouts.Client
	SelfActionName string
}

func (a *HangoutsAction) NotifyPullRequest(event *github.PullRequestEvent, filters ...PullRequestFilter) error {
	var title string
	switch *event.Action {
	case "opened":
		title = "New pull request is opened"
	case "reopened":
		title = "Pull request re-opened"
	case "synchronize":
		title = "Pull request updated"
	default:
		return nil
	}
	repo := *event.Repo.Name
	owner := *event.Repo.Owner.Login
	pr := event.PullRequest
	prKey := fmt.Sprintf("%s/%s-%d", owner, repo, *event.PullRequest.Number)
	for _, filter := range filters {
		if !filter(event) {
			return nil
		}
	}
	_, err := a.hangoutsClient.Send(prKey, &hangouts.Message{
		Cards: []*hangouts.Card{
			{
				Header: makeCardHeader(title, *pr.Title, StatusInProgress),
				Sections: []*hangouts.Section{
					makeViewSection(prKey, *pr.HTMLURL),
					makeAuthorSection(*pr.User.Login, *pr.User.HTMLURL, *pr.User.AvatarURL),
				},
			},
		},
	})
	return err
}

func (a *HangoutsAction) NotifyPullRequestChecks(event *github.PullRequestEvent, filters ...PullRequestChecksFilter) error {
	switch *event.Action {
	case "opened":
	case "reopened":
	case "synchronize":
	default:
		return nil
	}
	repo := *event.Repo.Name
	owner := *event.Repo.Owner.Login
	ref := *event.PullRequest.Head.SHA
	pr := event.PullRequest
	prKey := fmt.Sprintf("%s/%s-%d", owner, repo, *event.PullRequest.Number)
	// checks, err := a.GetChecks(context.Background(), "istio", "istio", "94f856ec6ccf5244a62e68c92d7f5dc23e0e4f09")
	checks, err := a.GetChecks(context.Background(), owner, repo, ref)
	if err != nil {
		return err
	}
	for _, filter := range filters {
		if !filter(event, checks) {
			return nil
		}
	}
	if checks.Empty() {
		return nil
	}
	overallStatus := checks.OverallStatus()
	var title string
	switch overallStatus {
	case StatusFailure:
		title = "Some checks were not successful"
	case StatusSuccess:
		title = "All checks have passed"
	default:
		title = "Checks are running"
	}
	_, err = a.hangoutsClient.Send(prKey, &hangouts.Message{
		Cards: []*hangouts.Card{
			{
				Header: makeCardHeader(title, *pr.Title, overallStatus),
				Sections: []*hangouts.Section{
					makeViewSection(prKey, *pr.HTMLURL),
					makeAuthorSection(*pr.User.Login, *pr.User.HTMLURL, *pr.User.AvatarURL),
					makeChecksSection(checks),
				},
			},
		},
	})
	return err
}

func (a *HangoutsAction) GetChecks(ctx context.Context, owner, repo, ref string) (Checks, error) {
	checks := make(Checks)
	statusList, _, err := a.githubClient.Repositories.ListStatuses(ctx, owner, repo, ref, nil)
	if err != nil {
		return checks, err
	}
	checksList, _, err := a.githubClient.Checks.ListCheckRunsForRef(ctx, owner, repo, ref, nil)
	if err != nil {
		return checks, err
	}
	log.Printf("%+v\n", checksList)
	for _, s := range statusList {
		status := statusFromGithubStatus(s)
		checks[status] = append(checks[status], Check{
			Status:    status,
			Name:      *s.Context,
			Message:   *s.Description,
			AvatarUrl: *s.Creator.AvatarURL,
			TargetUrl: *s.TargetURL,
		})
	}

	for _, c := range checksList.CheckRuns {
		// Don't need to include own check
		if *c.Name == a.SelfActionName {
			continue
		}
		status := statusFromGithubCheckRun(c)
		checks[status] = append(checks[status], Check{
			Status:    status,
			Name:      *c.Name,
			Message:   string(status),
			AvatarUrl: *c.App.Owner.AvatarURL,
			TargetUrl: *c.HTMLURL,
		})
	}
	return checks, nil
}

func imageFromStatus(s Status) string {
	switch s {
	case StatusSuccess:
		return ImageSuccess
	case StatusFailure:
		return ImageFailure
	case StatusInProgress:
		fallthrough
	default:
		return ImageInProgress
	}
}

func statusFromGithubCheckRun(c *github.CheckRun) Status {
	if *c.Status != "completed" {
		return StatusInProgress
	}
	if *c.Conclusion == "success" {
		return StatusSuccess
	}
	return StatusFailure
}

func statusFromGithubStatus(s *github.RepoStatus) Status {
	switch *s.State {
	case "success":
		return StatusSuccess
	case "failure":
		return StatusFailure
	case "pending":
		fallthrough
	default:
		return StatusInProgress
	}
}

func (c Checks) OverallStatus() Status {
	if _, ok := c[StatusFailure]; ok {
		return StatusFailure
	}
	if _, ok := c[StatusInProgress]; ok {
		return StatusInProgress
	}
	return StatusSuccess
}

func (c Checks) ToList() []Check {
	var checks []Check
	for _, v := range c {
		checks = append(checks, v...)
	}
	return checks
}

func (c Checks) Empty() bool {
	return len(c) == 0
}

func makeCardHeader(title, subTitle string, status Status) *hangouts.CardHeader {
	return &hangouts.CardHeader{
		Title:      title,
		Subtitle:   subTitle,
		ImageUrl:   imageFromStatus(status),
		ImageStyle: "IMAGE",
	}
}

func makeViewSection(content, url string) *hangouts.Section {
	return &hangouts.Section{
		Widgets: []*hangouts.WidgetMarkup{
			{
				// Buttons: []*hangouts.Button{
				// 	{
				// 		TextButton: &hangouts.TextButton{
				// 			Text: fmt.Sprintf("View %s #%d", repo, pr.Number),
				// 			OnClick: &hangouts.OnClick{
				// 				OpenLink: &hangouts.OpenLink{
				// 					Url: pr.HtmlUrl,
				// 				},
				// 			},
				// 		},
				// 	},
				// },
				KeyValue: &hangouts.KeyValue{
					Content: content,
					Button: &hangouts.Button{
						TextButton: &hangouts.TextButton{
							Text: "View",
							OnClick: &hangouts.OnClick{
								OpenLink: &hangouts.OpenLink{
									Url: url,
								},
							},
						},
					},
				},
			},
		},
	}
}

func makeAuthorSection(username, profileUrl, avatarUrl string) *hangouts.Section {
	return &hangouts.Section{
		Widgets: []*hangouts.WidgetMarkup{
			{
				KeyValue: &hangouts.KeyValue{
					IconUrl:  avatarUrl,
					TopLabel: "Author",
					Content:  username,
					Button: &hangouts.Button{
						ImageButton: &hangouts.ImageButton{
							IconUrl: ImageGitHubAvatar,
							Name:    "View Profile",
							OnClick: &hangouts.OnClick{
								OpenLink: &hangouts.OpenLink{
									Url: profileUrl,
								},
							},
						},
					},
				},
			},
		},
	}
}

func makeChecksSection(checks Checks) *hangouts.Section {
	var checksWidgets []*hangouts.WidgetMarkup
	for _, v := range checks.ToList() {
		checksWidgets = append(checksWidgets, &hangouts.WidgetMarkup{
			KeyValue: &hangouts.KeyValue{
				IconUrl:  imageFromStatus(v.Status),
				Content:  v.Message,
				TopLabel: v.Name,
				Button: func() *hangouts.Button {
					if len(v.TargetUrl) > 0 {
						return &hangouts.Button{
							ImageButton: &hangouts.ImageButton{
								IconUrl: v.AvatarUrl,
								Name:    "View",
								OnClick: &hangouts.OnClick{
									OpenLink: &hangouts.OpenLink{
										Url: v.TargetUrl,
									},
								},
							},
						}
					}
					return nil
				}(),
			},
		})
	}
	return &hangouts.Section{
		Header:  "Checks",
		Widgets: checksWidgets,
	}
}
