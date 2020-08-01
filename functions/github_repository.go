package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
	"time"
)

type PullRequest struct {
	Url       string
	Number    int
	Title     string
	CreatedAt time.Time
}

func PullRequests(repositoryName string) (pullRequests []PullRequest) {
	client := githubClient()

	opts := &github.PullRequestListOptions{"open", "", "", "created", "desc", github.ListOptions{PerPage: 15}}
	githubPullRequests, _, err := client.PullRequests.List(context.Background(), "vinomofo", repositoryName, opts)

	if err != nil {
		fmt.Printf("Problem in getting pull requests: %v", err)
		os.Exit(1)
	}

	for _, value := range githubPullRequests {
		pullRequest := buildPullRequest(*value)
		pullRequests = append(pullRequests, pullRequest)
	}
	
	return pullRequests
}

func githubClient() (client *github.Client) {
	githubToken 	:= os.Getenv("GITHUB_TOKEN")
	tokenService	:= oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	tokenClient 	:= oauth2.NewClient(context.Background(), tokenService)
	
	return github.NewClient(tokenClient)
}

func buildPullRequest(githubPullRequest github.PullRequest) PullRequest {
	return PullRequest{
		Url:       *githubPullRequest.URL,
		Number:    *githubPullRequest.Number,
		Title:     *githubPullRequest.Title,
		CreatedAt: *githubPullRequest.CreatedAt,
	}
}
