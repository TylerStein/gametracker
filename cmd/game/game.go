package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/TylerStein/gametracker/internal/data"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/google/uuid"
)

// Define a dependencies struct so the handler can be easily called with mock providers
type Handler struct {
	Database   dynamodbiface.DynamoDBAPI
	GamesTable string
}

// Make the request handler a Handler reciever so the dependencies can be accessed while respecting the AWS function signature
func (d *Handler) handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body data.Game
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		fmt.Println("Error calling json.Unmarshal on request body")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Unexpected body"),
			StatusCode: 400,
		}, nil
	}

	bodyErr := body.Validate()
	if bodyErr != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Malformed body: %v", bodyErr.Error()),
			StatusCode: 400,
		}, nil
	}

	body.Id = uuid.NewString()
	item, err := dynamodbattribute.MarshalMap(body)
	if err != nil {
		fmt.Println("Could not call dynamodbattribute.MarshalMap on body")
		fmt.Println(body)
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Internal server error"),
			StatusCode: 500,
		}, nil
	}

	input := dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(d.GamesTable),
	}

	_, err = d.Database.PutItem(&input)
	if err != nil {
		fmt.Println("Could not call d.Database.PutItem on input")
		fmt.Println(body)
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Internal server error"),
			StatusCode: 500,
		}, nil
	}

	responseBytes, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Could not marshal body to json")
		fmt.Println(body)
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Internal server error"),
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	sess := session.Must(session.NewSession())
	handler := Handler{
		Database:   dynamodb.New(sess),
		GamesTable: os.Getenv("GAME_TABLE_NAME"),
	}

	lambda.Start(handler.handleRequest)
}
