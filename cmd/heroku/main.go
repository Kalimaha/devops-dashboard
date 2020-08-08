package main

import (
	"encoding/json"
	"github.com/Kalimaha/devops-dashboard/pkg/repositories"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func herokuHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	repositoryName, _ := request.QueryStringParameters["repositoryName"]

	releases := repositories.ListReleasesFor(repositoryName)
	commits := repositories.CompareCommits(repositoryName, releases[1].CommitID, releases[0].CommitID)

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

func main() {
	lambda.Start(herokuHandler)
}
