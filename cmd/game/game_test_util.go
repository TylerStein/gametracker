package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

	"github.com/TylerStein/gametracker/internal/data"
)

type mockedPutItem struct {
	dynamodbiface.DynamoDBAPI
	Response dynamodb.PutItemOutput
}

func (d mockedPutItem) PutItem(in *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &d.Response, nil
}

func createMockPostHandler() Handler {
	return Handler{
		Database:   mockedPutItem{Response: dynamodb.PutItemOutput{}},
		GamesTable: "Games",
	}
}

func createMockRequestWithBody(body string) events.APIGatewayProxyRequest {
	return events.APIGatewayProxyRequest{
		Resource:                        "",
		Path:                            "",
		HTTPMethod:                      "",
		Headers:                         map[string]string{},
		MultiValueHeaders:               map[string][]string{},
		QueryStringParameters:           map[string]string{},
		MultiValueQueryStringParameters: map[string][]string{},
		PathParameters:                  map[string]string{},
		StageVariables:                  map[string]string{},
		RequestContext:                  events.APIGatewayProxyRequestContext{},
		Body:                            body,
		IsBase64Encoded:                 false,
	}
}

func createNewGameBody() string {
	game := data.Game{
		Name: "Half-Life",
	}

	bytes, err := json.Marshal(game)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
