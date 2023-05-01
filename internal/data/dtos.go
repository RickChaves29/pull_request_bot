package data

type GithubUser struct {
	UserName   string `json:"login"`
	UserAvatar string `json:"avatar_url"`
}
type PullRequest struct {
	Url         string     `json:"html_url"`
	State       string     `json:"state"`
	Title       string     `json:"title"`
	User        GithubUser `json:"user"`
	Description string     `json:"body"`
	DatePR      string     `json:"created_at"`
}
