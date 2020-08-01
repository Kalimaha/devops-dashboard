package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	repositoryName, _ := request.QueryStringParameters["repositoryName"]

	pullRequests := PullRequests(repositoryName)
	body, _ := json.Marshal(pullRequests)
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
	lambda.Start(handler)
}

func Sum(a int, b int) int {
	return a + b
}
