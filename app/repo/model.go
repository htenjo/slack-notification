package repo

import "time"

type PrReviewFlag string

const (
	Green  PrReviewFlag = "Green"
	Yellow PrReviewFlag = "Yellow"
	Red    PrReviewFlag = "Red"
)

type PullRequestDetail struct {
	Url        string       `json:"url"`
	HtmlUrl    string       `json:"html_url"`
	Title      string       `json:"title"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
	ReviewFlag PrReviewFlag `json:"review_flag"`
}
