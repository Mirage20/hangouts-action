package main

import (
	"fmt"

	"github.com/mirage20/hangouts-action/github"
	"github.com/mirage20/hangouts-action/hangouts"
)

const (
	ImageSuccess      = "https://www.shareicon.net/download/2017/02/09/878601_check_512x512.png"
	ImageFailure      = "https://www.shareicon.net/download/2017/02/09/878603_close_512x512.png"
	ImageInProgress   = "https://www.shareicon.net/download/2017/02/09/878594_gear_512x512.png"
	ImageGitHubAvatar = "https://avatars0.githubusercontent.com/in/15368?s=40&v=4"
)

type Status string

const (
	StatusSuccess    Status = "Success"
	StatusFailure    Status = "Failure"
	StatusInProgress Status = "InProgress"
)

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
	if c.Status != "completed" {
		return StatusInProgress
	}
	if c.Conclusion == "success" {
		return StatusSuccess
	}
	return StatusFailure
}

func statusFromGithubStatus(s *github.Status) Status {
	switch s.State {
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

func makeMessageFromPullRequest(pr *github.PullRequest, title string, repo string, status *github.StatusResponse, checks *github.CheckRunsResponse) *hangouts.Message {

	statusSection := &hangouts.Section{
		Header: "Checks",
	}
	statusMap := make(map[Status]bool)
	var statusWidgets []*hangouts.WidgetMarkup
	for _, s := range status.Statuses {
		status := statusFromGithubStatus(&s)
		statusMap[status] = true
		statusWidgets = append(statusWidgets, &hangouts.WidgetMarkup{
			KeyValue: &hangouts.KeyValue{
				IconUrl:  imageFromStatus(status),
				Content:  s.Description,
				TopLabel: s.Context,
				Button: func() *hangouts.Button {
					if len(s.TargetUrl) > 0 {
						return &hangouts.Button{
							ImageButton: &hangouts.ImageButton{
								IconUrl: s.AvatarUrl,
								Name:    "View",
								OnClick: &hangouts.OnClick{
									OpenLink: &hangouts.OpenLink{
										Url: s.TargetUrl,
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

	for _, c := range checks.CheckRuns {
		// Don't need to include own check
		if c.Name == "Hangouts" {
			continue
		}
		status := statusFromGithubCheckRun(&c)
		statusMap[status] = true
		statusWidgets = append(statusWidgets, &hangouts.WidgetMarkup{
			KeyValue: &hangouts.KeyValue{
				IconUrl:  imageFromStatus(status),
				Content:  string(status),
				TopLabel: c.Name,
				Button: func() *hangouts.Button {
					if len(c.HtmlUrl) > 0 {
						return &hangouts.Button{
							ImageButton: &hangouts.ImageButton{
								IconUrl: c.App.Owner.AvatarUrl,
								Name:    "View",
								OnClick: &hangouts.OnClick{
									OpenLink: &hangouts.OpenLink{
										Url: c.HtmlUrl,
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

	statusSection.Widgets = statusWidgets

	return &hangouts.Message{
		Cards: []*hangouts.Card{
			{
				Header: makeCardHeader(pr, title, getOverallStatus(statusMap)),
				Sections: []*hangouts.Section{
					makeViewPullRequestSection(pr, repo),
					makeAuthorSection(pr),
					statusSection,
				},
			},
		},
	}
}

func getOverallStatus(statusMap map[Status]bool) Status {
	if _, ok := statusMap[StatusFailure]; ok {
		return StatusFailure
	}
	if _, ok := statusMap[StatusInProgress]; ok {
		return StatusInProgress
	}
	return StatusSuccess
}

func makeCardHeader(pr *github.PullRequest, title string, overallStatus Status) *hangouts.CardHeader {
	return &hangouts.CardHeader{
		Title:      title,
		Subtitle:   pr.Title,
		ImageUrl:   imageFromStatus(overallStatus),
		ImageStyle: "IMAGE",
	}
}

func makeViewPullRequestSection(pr *github.PullRequest, repo string) *hangouts.Section {
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
					Content: fmt.Sprintf("%s#%d", repo, pr.Number),
					Button: &hangouts.Button{
						TextButton: &hangouts.TextButton{
							Text: "View",
							OnClick: &hangouts.OnClick{
								OpenLink: &hangouts.OpenLink{
									Url: pr.HtmlUrl,
								},
							},
						},
					},
				},
			},
		},
	}
}

func makeAuthorSection(pr *github.PullRequest) *hangouts.Section {
	return &hangouts.Section{
		Widgets: []*hangouts.WidgetMarkup{
			{
				KeyValue: &hangouts.KeyValue{
					IconUrl:  pr.User.AvatarUrl,
					TopLabel: "Author",
					Content:  pr.User.LoginName,
					Button: &hangouts.Button{
						ImageButton: &hangouts.ImageButton{
							IconUrl: ImageGitHubAvatar,
							Name:    "View Profile",
							OnClick: &hangouts.OnClick{
								OpenLink: &hangouts.OpenLink{
									Url: pr.User.HtmlUrl,
								},
							},
						},
					},
				},
			},
		},
	}
}
