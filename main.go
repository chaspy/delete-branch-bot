package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
)

type Response struct {
	Deleted bool
	Repo    string
	Branch  string
}

type GithubKeys struct {
	AccessToken string `json:"GITHUB_ACCESS_TOKEN"`
}

func delete_branch(event PullRequestEvent) (Response, error) {
	var errorType error
	res := Response{Repo: event.PullRequest.Head.Repo.FullName, Branch: event.PullRequest.Head.Ref}

	if event.Action == "closed" || event.Action == "merged" {
		fmt.Println(event.Action)

		// Create client via go-github library
		tr := http.DefaultTransport
		GITHUB_APP_ID, _ := strconv.Atoi(os.Getenv("GITHUB_APP_IDENTIFIER"))
		itr, err := ghinstallation.NewKeyFromFile(tr, GITHUB_APP_ID, event.Installation.ID, os.Getenv("GITHUB_PRIVATE_KEY"))
		if err != nil {
			log.Fatal(err)
			res.Deleted = false
			return res, errorType
		}

		fmt.Printf("itr:%v", itr)
		// Use installation transport with github.com/google/go-github
		client := github.NewClient(&http.Client{Transport: itr})

		//Create request
		Owner := event.PullRequest.Head.Repo.Owner.Login
		RepoName := event.PullRequest.Head.Repo.Name
		HeadRefs := "heads/" + event.PullRequest.Head.Ref

		ctx := context.Background()
		resp, err := client.Git.DeleteRef(ctx, Owner, RepoName, HeadRefs)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error Occured")
			res.Deleted = false
			return res, errorType
		}
		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", resp.Body)

		res.Deleted = true
	} else {
		fmt.Println("no match action")
		res.Deleted = false
	}

	return res, nil
}

func main() {
	lambda.Start(delete_branch)
}
