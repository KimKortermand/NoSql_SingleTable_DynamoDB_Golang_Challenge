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

	out, err := dynamoClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(r.Table), //("kk-test2"),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: r.DeviceGroupID}, // "dg#3"},
			"SK": &types.AttributeValueMemberS{Value: r.DeviceGroupID}, // "dg#3"},
		},
	})
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}

	type Item struct {
		PK         string
		SK         string
		EntityType string
		Name       string
	}
	var item Item
	err = attributevalue.UnmarshalMap(out.Item, &item)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(item.EntityType)
	log.Println(item.Name)

	response, err := json.Marshal(map[string]string{
		"entitytype": item.EntityType,
		"name":       item.Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	return Response{Body: string(response), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
