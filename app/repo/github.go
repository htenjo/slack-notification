package repo

import (
	"encoding/json"
	"github.com/htenjo/slack-notification/app/http"
	"github.com/spf13/viper"
	"time"
)

const (
	prNewHours            = 4
	prMiddleHours         = 8
	repositoryUrlListKey  = "reposUrl"
	repositoryUsernameKey = "repoUsername"
	repositoryPasswordKey = "repoPassword"
	pullUrlFilter         = "/pulls?state=open&sort=updated"
)

func OpenPullRequests() *[]PullRequestDetail {
	repoUrl := viper.GetStringSlice(repositoryUrlListKey)
	repoUsername := viper.GetString(repositoryUsernameKey)
	repoPwd := viper.GetString(repositoryPasswordKey)
	var pullRequestResponse []PullRequestDetail
	var details []PullRequestDetail

	for _, repoUrl := range repoUrl {
		repoUrl += pullUrlFilter
		pullRequestJson := http.GetWithBasicAuth(repoUrl, repoUsername, repoPwd)
		_ = json.Unmarshal([]byte(pullRequestJson), &details)
		pullRequestResponse = append(pullRequestResponse, processPrOpen(details)...)
	}

	return &pullRequestResponse
}

func processPrOpen(pullRequests []PullRequestDetail) []PullRequestDetail {
	currentTime := time.Now()

	for i := 0; i < len(pullRequests); i++ {
		pr := &(pullRequests)[i]
		openHours := currentTime.Sub(pr.CreatedAt).Hours()

		if openHours < prNewHours {
			pr.ReviewFlag = Green
		} else if prNewHours <= openHours && openHours < prMiddleHours {
			pr.ReviewFlag = Yellow
		} else {
			pr.ReviewFlag = Red
		}
	}

	return pullRequests
}
