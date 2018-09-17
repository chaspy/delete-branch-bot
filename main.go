package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

type PullRequestEvent struct {
	Action string `json:"action"`
	Number int    `json:"number"`
}

func delete_branch(event PullRequestEvent) (MyResponse, error) {
	fmt.Println("log start")
	fmt.Println(event)
	fmt.Println("log end")
	return MyResponse{Message: fmt.Sprintf("PullRequest action is %s!!", event.Action)}, nil
}

func main() {
	lambda.Start(delete_branch)
}
