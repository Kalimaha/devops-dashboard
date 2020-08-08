package repositories

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
	UpdatedAt  time.Time
	Reviews    []Review
	AuthorName string
	AuthorURL  string
	AvatarURL  string
}

type Review struct {
	ReviewerURL  string
	ReviewerName string
	State        string
	SubmittedAt  time.Time
}

type Commit struct {
	AuthorLogin string
	AuthorURL   string
	Message     string
}

func CompareCommits(repositoryName string, head string, tail string) (commits []Commit) {
	fmt.Printf("\tGH-REPO: repositoryName -> %s\n", repositoryName)
	fmt.Printf("\tGH-REPO: repositoryName -> %s\n", head)
	fmt.Printf("\tGH-REPO: repositoryName -> %s\n", tail)
	// client := githubClient()
	// commitsComparison, _, _ := client.Repositories.CompareCommits(context.Background(), "vinomofo", repositoryName, head, tail)

	// fmt.Printf("\tGH-REPO: commits -> %v\n", len(commitsComparison.Commits))
	// for _, commit := range commitsComparison.Commits {
	// 	fmt.Printf("\tGH-REPO: CONVERT -> %v\n", commit)
	// 	commits = append(commits, buildCommit(*commit))
	// }

	return commits
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
	fmt.Printf("\tgithubClient -> githubToken? %v\n", githubToken)
	
	tokenService := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken})
	fmt.Printf("\tgithubClient -> tokenService? %v\n", tokenService)
	
	tokenClient := oauth2.NewClient(context.Background(), tokenService)
	fmt.Printf("\tgithubClient -> tokenClient? %v\n", tokenClient)

	return github.NewClient(tokenClient)
}

func buildPullRequest(githubPullRequest github.PullRequest, reviews []Review) PullRequest {
	return PullRequest{
		Url:        *githubPullRequest.HTMLURL,
		Number:     *githubPullRequest.Number,
		Title:      *githubPullRequest.Title,
		CreatedAt:  *githubPullRequest.CreatedAt,
		UpdatedAt:  *githubPullRequest.UpdatedAt,
		Reviews:    reviews,
		AuthorName: *githubPullRequest.User.Login,
		AuthorURL:  *githubPullRequest.User.Login,
		AvatarURL:  *githubPullRequest.User.AvatarURL,
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

func buildCommit(commit github.RepositoryCommit) Commit {
	fmt.Printf("\tCOMMIT %+v\n", commit)
	return Commit{
		Message:     *commit.Commit.Message,
		AuthorLogin: *commit.Author.Login,
		AuthorURL:   *commit.Author.HTMLURL,
	}
}
