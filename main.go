package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type MyResponse struct {
	Message string `json:"Answer:"`
}

type PullRequestEvent struct {
	Action      string `json:"action"`
	PullRequest struct {
		Head struct {
			Ref  string `json:"ref"`
			Sha  string `json:"sha"`
			Repo struct {
				ID       int    `json:"id"`
				NodeID   string `json:"node_id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Owner    struct {
					Login    string `json:"login"`
					ID       int    `json:"id"`
					NodeID   string `json:"node_id"`
					ReposURL string `json:"repos_url"`
				} `json:"owner"`
			} `json:"repo"`
		} `json:"head"`
	}
}

type GithubKeys struct {
	AccessToken string `json:"GITHUB_ACCESS_TOKEN"`
}

func delete_branch(event PullRequestEvent) (MyResponse, error) {
	fmt.Println("log start")
	fmt.Println(event)
	fmt.Println("log end")

	var message MyResponse
	var errorType error

	// hi @kamontia! I checked this lambda function works from GitHub-App via API-Gateway.
	// I leave the hint to implement we realy want to do!
	//
	// currently, I already cache PullRequestEvent from GitHub,
	// and Got values we need.
	//
	// action: event.Action (i.e closed or merged)
	// repo  : event.PullRequest.Head.Repo.FullName (i.e chaspy/delete-branch)
	// branch: event.PullRequest.Head.Ref (i.e test-branch)
	// ref https://github.com/chaspy/test-delete-branch/pull/1
	// (I used this PR to test)
	//
	// So what you should implement is very simple!
	//
	// if action is "closed" or "merged" then
	//  delete branch(repo,branch)
	// end
	// return nothing(204 No Content)
	//
	// If there is concern, it will be part of certification.
	// ref: https://github.com/github/platform-samples/blob/b047a807dd43a3f76c2cbf0e85af3ffadeb2b880/api/ruby/building-your-first-github-app/advanced_server.rb#L24-L44
	if event.Action == "closed" || event.Action == "merged" {
		fmt.Println(event.Action)
		var keys GithubKeys
		raw, err := ioutil.ReadFile("./keys.json")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		json.Unmarshal(raw, &keys)

		// Authentication
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: keys.AccessToken},
		)
		tc := oauth2.NewClient(ctx, ts)

		// Create client via go-github library
		client := github.NewClient(tc)

		//Create request
		Owner := event.PullRequest.Head.Repo.Owner.Login
		RepoName := event.PullRequest.Head.Repo.Name
		HeadRefs := event.PullRequest.Head.Ref

		resp, err := client.Git.DeleteRef(ctx, Owner, RepoName, HeadRefs)
		if err != nil {
			fmt.Println(err)
			return message, errorType
		}
		fmt.Println("response Status : ", resp.Status)
		fmt.Println("response Headers : ", resp.Header)
		fmt.Println("response Body : ", resp.Body)

	} else {
		fmt.Println("no match action")
	}

	return MyResponse{Message: fmt.Sprintf("PullRequest action is %s!! repo is %s!! branch name is %s!!", event.Action, event.PullRequest.Head.Ref, event.PullRequest.Head.Repo.FullName)}, nil
}

func main() {
	lambda.Start(delete_branch)
}
