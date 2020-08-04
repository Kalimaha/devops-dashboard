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
	Url        string
	Number     int
	Title      string
	CreatedAt  time.Time
	Reviews    []Review
	AuthorName string
	AuthorURL  string
}

type Review struct {
	ReviewerURL  string
	ReviewerName string
	State        string
	SubmittedAt  time.Time
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
		reviews := fetchReviews(client, repositoryName, *value.Number)
		pullRequest := buildPullRequest(*value, reviews)
		pullRequests = append(pullRequests, pullRequest)
	}

	return pullRequests
}

func fetchReviews(client *github.Client, repositoryName string, pullRequestNumber int) []Review {
	reviews := make([]Review, 0)
	opt := &github.ListOptions{PerPage: 10}
	githubReviews, _, err := client.PullRequests.ListReviews(context.Background(), "vinomofo", repositoryName, pullRequestNumber, opt)

	if err != nil {
		fmt.Printf("Problem in getting reviews: %v", err)
		os.Exit(1)
	}

	for _, githubReview := range githubReviews {
		review := buildReview(*githubReview)
		reviews = append(reviews, review)
	}

	return reviews
}

func githubClient() (client *github.Client) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	tokenService := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	tokenClient := oauth2.NewClient(context.Background(), tokenService)

	return github.NewClient(tokenClient)
}

func buildPullRequest(githubPullRequest github.PullRequest, reviews []Review) PullRequest {
	return PullRequest{
		Url:        *githubPullRequest.HTMLURL,
		Number:     *githubPullRequest.Number,
		Title:      *githubPullRequest.Title,
		CreatedAt:  *githubPullRequest.CreatedAt,
		Reviews:    reviews,
		AuthorName: *githubPullRequest.User.Login,
		AuthorURL:  *githubPullRequest.User.Login,
	}
}

func buildReview(githubPullRequestReview github.PullRequestReview) Review {
	return Review{
		ReviewerURL:  *githubPullRequestReview.User.HTMLURL,
		ReviewerName: *githubPullRequestReview.User.Login,
		State:        *githubPullRequestReview.State,
		SubmittedAt:  *githubPullRequestReview.SubmittedAt,
	}
}
