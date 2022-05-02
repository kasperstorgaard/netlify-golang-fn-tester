package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "hi",
		// this header shows up in debugging/devtools
		Headers: map[string]string{
			"X-Foo": "1",
		},
		// this doesn't seem to work
		MultiValueHeaders: map[string][]string{
			"X-Bar": {"2"},
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
