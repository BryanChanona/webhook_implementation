package value_objects


type PullRequestEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
	Repository  Repository  `json:"repository"`
}

type PullRequest struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	User    User   `json:"user"`
	Head    Branch `json:"head"`
	Base    Branch `json:"base"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
	Merged  bool   `json:"merged"`
}

type Branch struct {
	Ref string `json:"ref"`
	Sha string `json:"sha"`
}



type Repository struct {
	FullName string `json:"full_name"`
	URL      string `json:"url"`
}

type User struct {
	Login string `json:"login"`
}
