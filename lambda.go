package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/chaspy/delete-branch/model"
)

type MyEvent struct {
	Name string `json:"What is your name?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
}

func hello(event MyEvent) (MyResponse, error) {
  data := new(model.PullRequestEvent)
  var _ = data // avoid build error
	return MyResponse{Message: fmt.Sprintf("Hello %s!!", event.Name)}, nil
}

func main() {
	lambda.Start(hello)
}
