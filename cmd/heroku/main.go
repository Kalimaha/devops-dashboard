package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func herokuHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	repositoryName, _ := request.QueryStringParameters["repositoryName"]

	releases := ListReleasesFor(repositoryName)
	commits := CompareCommits(repositoryName, releases[1].CommitID, releases[0].CommitID)

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
