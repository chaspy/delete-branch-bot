package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chaspy/delete-branch/model"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

func delete_branch(event model.PullRequestEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("PullRequest action is %s!!", event.Payload.Action)}, nil
}

func main() {
	lambda.Start(delete_branch)
}
