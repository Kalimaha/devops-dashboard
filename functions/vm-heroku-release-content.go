package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	repositoryName, _ := request.QueryStringParameters["repositoryName"]

	releases := ListReleasesFor(repositoryName)
	commits := CompareCommits(repositoryName, releases[1], releases[0])

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
}

func main() {
	lambda.Start(handler)
}
