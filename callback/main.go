package main

import (
	"callback/models"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
Handler should return http status:
200 - OK
401 - Unauthorized
404 - Not Found
500 - Internal Server Error
502 - Bad Gateway
*/

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// parse request body
	var callback models.CatalogImportRequest
	err := json.Unmarshal([]byte(request.Body), &callback)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(handler)
}
