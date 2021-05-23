package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

type verb string

const (
	get verb = "GET"
)

func GetWithBasicAuth(url, username, passwd string) string {
	return requestBasicAuth(url, username, passwd, get)
}

func requestBasicAuth(url, username, passwd string, verb verb) string {
	client := &http.Client{}
	req, err := http.NewRequest(string(verb), url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	s := string(bodyText)
	return s
}
