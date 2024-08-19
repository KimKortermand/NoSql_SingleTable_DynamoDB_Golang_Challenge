package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.LambdaFunctionURLRequest) (Response, error) {

	log.Println("Hello World 3")

	response := "Hello World 3!"
	return Response{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
