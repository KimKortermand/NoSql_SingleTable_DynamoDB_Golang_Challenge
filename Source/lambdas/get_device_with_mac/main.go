package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context, request events.LambdaFunctionURLRequest) (Response, error) {

	type RequestStruct struct {
		Table string `json:"table"`
		MAC   string `json:"mac"`
	}
	var r RequestStruct

	body_str := request.Body
	body_bytes := []byte(body_str)
	err := json.Unmarshal(body_bytes, &r)
	if err != nil {
		log.Fatalf("unmarshalling json: %s, %v", body_str, err)
	}
	log.Println(r.Table)
	log.Println(r.MAC)

	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-central-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	dynamoClient := dynamodb.NewFromConfig(cfg)

	out, err := dynamoClient.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(r.Table),
		IndexName:              aws.String("EntityType-MAC-index"),
		KeyConditionExpression: aws.String("EntityType = :pk AND MAC = :sk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{Value: "device"},
			":sk": &types.AttributeValueMemberS{Value: r.MAC},
		},
	})
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}

	log.Println(out.Items)

	type Item struct {
		PK         string `json:"pk"`
		SK         string `json:"sk"`
		EntityType string `json:"entitytype"`
		Name       string `json:"name"`
		MAC        string `json:"mac"`
		Status     string `json:"status"`
	}
	var items = []Item{}
	err = attributevalue.UnmarshalListOfMaps(out.Items, &items)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(items)

	response, err := json.Marshal(items)
	log.Println(string(response))

	return Response{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
