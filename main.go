package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

type PullRequestEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Head struct {
			Ref  string `json:"ref"`
			Repo struct {
				FullName string `json:"full_name"`
			} `json:"repo"`
		} `json:"head"`
	} `json:"pull_request"`
}

func delete_branch(event PullRequestEvent) (MyResponse, error) {
	fmt.Println("log start")
	fmt.Println(event)
	fmt.Println("log end")
	return MyResponse{Message: fmt.Sprintf("PullRequest action is %s!! repo is %s!! branch name is %s!!", event.Action, event.PullRequest.Head.Ref, event.PullRequest.Head.Repo.FullName)}, nil
}

func main() {
	lambda.Start(delete_branch)
}
