package main

import (
	"encoding/json"
	"fmt"
	"github.com/Kalimaha/devops-dashboard/pkg/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func herokuHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	githubRepositoryName, _ := request.QueryStringParameters["repositoryName"]
	fmt.Printf("githubRepositoryName %s\n", githubRepositoryName)

	herokuRepositoryName := herokuRepositoryName(githubRepositoryName)
	fmt.Printf("herokuRepositoryName %s\n", herokuRepositoryName)

	releases := repositories.ListReleasesFor(herokuRepositoryName)
	fmt.Printf("RELEASES %+v\n", releases)
	fmt.Printf("RELEASES %+v\n", len(releases))
	fmt.Printf("RELEASES 0 %+v\n", releases[0])
	fmt.Printf("RELEASES 0 %+v\n", releases[0].CommitID)
	fmt.Printf("RELEASES 1 %+v\n", releases[1])
	fmt.Printf("RELEASES 1 %+v\n", releases[1].CommitID)

	commits := repositories.CompareCommits(githubRepositoryName, releases[1].CommitID, releases[0].CommitID)
	fmt.Printf("COMMITS %+v\n", commits)

	body, _ := json.Marshal(commits)
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	}, nil
}

func herokuRepositoryName(githubRepositoryName string) string {
	if githubRepositoryName == "vinomofo" {
		return "vinomofo-au"
	} else if githubRepositoryName == "vino-subscription" {
		return "vino-subscription-au"
	} else if githubRepositoryName == "vino-warehouse" {
		return "production-vino-warehouse"
	} else {
		return githubRepositoryName
	}
}

func main() {
	lambda.Start(herokuHandler)
}
