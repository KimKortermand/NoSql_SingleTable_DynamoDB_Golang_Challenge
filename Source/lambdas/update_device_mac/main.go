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
		Table         string `json:"table"`
		DeviceGroupID string `json:"devicegroupid"`
		DeviceID      string `json:"deviceid"`
		MAC           string `json:"mac"`
	}
	var r RequestStruct

	body_str := request.Body
	body_bytes := []byte(body_str)
	err := json.Unmarshal(body_bytes, &r)
	if err != nil {
		log.Fatalf("unmarshalling json: %s, %v", body_str, err)
	}
	log.Println(r.Table)
	log.Println(r.DeviceGroupID)
	log.Println(r.DeviceID)
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

	type Item struct {
		PK string `json:"pk"`
		SK string `json:"sk"`
	}
	item := Item{
		PK: r.DeviceGroupID,
		SK: r.DeviceID,
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling item: %s", err)
	}

	//
	// For more modern impl see:
	//   https://docs.aws.amazon.com/code-library/latest/ug/go_2_dynamodb_code_examples.html
	//
	input := &dynamodb.UpdateItemInput{
		TableName:        aws.String(r.Table),
		Key:              av,
		UpdateExpression: aws.String("SET MAC = :mac"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":mac": &types.AttributeValueMemberS{Value: r.MAC},
		},
	}

	_, err = dynamoClient.UpdateItem(context.TODO(), input)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}

	response := "Done!"
	return Response{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
