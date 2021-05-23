package main

import (
	"encoding/json"
	"fmt"
	"github.com/htenjo/slack-notification/app/config"
	"github.com/htenjo/slack-notification/app/repo"
)

func main() {
	config.Viper()
	openPRs := repo.OpenPullRequests()

	byteResponse, _ := json.MarshalIndent(openPRs, "", "  ")
	fmt.Printf("%v", string(byteResponse))
}
