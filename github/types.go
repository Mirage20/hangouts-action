package github

type ActionEvent struct {
	EventName string `json:"event_name,omitempty"`
	Event     Event  `json:"event,omitempty"`
}

type Event struct {
	Action      string      `json:"action,omitempty"`
	PullRequest PullRequest `json:"pull_request,omitempty"`
}

type PullRequest struct {
	Title   string `json:"title,omitempty"`
	Body    string `json:"body,omitempty"`
	HtmlUrl string `json:"html_url,omitempty"`
	Number  int    `json:"number,omitempty"`
	User    User   `json:"user,omitempty"`
	Head    Head   `json:"head,omitempty"`
}

type User struct {
	AvatarUrl string `json:"avatar_url,omitempty"`
	HtmlUrl   string `json:"html_url,omitempty"`
	LoginName string `json:"login,omitempty"`
}
type Head struct {
	Sha string `json:"sha,omitempty"`
}

type CheckRunsResponse struct {
	TotalCount int        `json:"total_count,omitempty"`
	CheckRuns  []CheckRun `json:"check_runs,omitempty"`
}

type CheckRun struct {
	Status     string      `json:"status,omitempty"`
	Conclusion string      `json:"conclusion,omitempty"`
	Name       string      `json:"name,omitempty"`
	HtmlUrl    string      `json:"html_url,omitempty"`
	App        CheckRunApp `json:"app,omitempty"`
}

type CheckRunApp struct {
	Description string           `json:"description,omitempty"`
	Name        string           `json:"name,omitempty"`
	Owner       CheckRunAppOwner `json:"owner,omitempty"`
}

type CheckRunAppOwner struct {
	AvatarUrl string `json:"avatar_url,omitempty"`
}

type StatusResponse struct {
	State    string   `json:"state,omitempty"`
	Statuses []Status `json:"statuses,omitempty"`
}

type Status struct {
	State       string `json:"state,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	Description string `json:"description,omitempty"`
	TargetUrl   string `json:"target_url,omitempty"`
	Context     string `json:"context,omitempty"`
}
