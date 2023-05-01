package data

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func ConnectionOnGithub(Url string) []PullRequest {
	res, _ := http.Get(Url)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("LOG [error-readall]: %v", err)
	}
	var data []PullRequest
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("LOG [error-json]: %v", err)
	}
	return data
}
