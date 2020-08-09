package repositories

import (
	"fmt"
	"log"
	"context"
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
	client := githubClient()
	fmt.Printf("COPARE %s AND %s FOR %s\n", head, tail, repositoryName)
	log.Printf("COPARE %s AND %s FOR %s\n", head, tail, repositoryName)
	commitsComparison, _, _ := client.Repositories.CompareCommits(context.Background(), "vinomofo", repositoryName, head, tail)

	for _, commit := range commitsComparison.Commits {
		fmt.Printf("ORIGINAL COMMIT: %+v\n", commit)
		log.Printf("ORIGINAL COMMIT: %+v\n", commit)
		builtCommit := buildCommit(*commit)
		fmt.Printf("BUILT COMMIT: %+v\n", builtCommit)
		log.Printf("BUILT COMMIT: %+v\n", builtCommit)
		commits = append(commits, builtCommit)
	}

	return commits
}

func PullRequests(repositoryName string) (pullRequests []PullRequest) {
	client := githubClient()

	opts := &github.PullRequestListOptions{"open", "", "", "created", "desc", github.ListOptions{PerPage: 15}}
	githubPullRequests, _, _ := client.PullRequests.List(context.Background(), "vinomofo", repositoryName, opts)

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
	githubReviews, _, _ := client.PullRequests.ListReviews(context.Background(), "vinomofo", repositoryName, pullRequestNumber, opt)

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
	return Commit{
		Message:     message(commit),
		// AuthorLogin: authorLogin(commit),
		// AuthorURL:   authorURL(commit),
		AuthorLogin: "",
		AuthorURL:   "",
	}
}

func message(commit github.RepositoryCommit) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = ""
		}
	}()

	output = *commit.Commit.Message
	return output
}

func authorLogin(commit github.RepositoryCommit) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = ""
		}
	}()

	output = *commit.Author.Login
	return output
}

func authorURL(commit github.RepositoryCommit) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = ""
		}
	}()

	output = *commit.Author.HTMLURL
	return output
}
