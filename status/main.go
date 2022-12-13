package main

import (
	"encoding/json"
	"fmt"
	orderstatus "status/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// TODO: remoteId and remoteOrderId are not used will be used
// TODO : logging should be added

/*
Handler should return http status:
200 - OK
401 - Unauthorized
404 - Not Found
500 - Internal Server Error
502 - Bad Gateway
*/

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// remoteId := request.PathParameters["remoteId"]
	// remoteOrderId := request.PathParameters["remoteOrderId"]

	// parse request body
	var orderStatus orderstatus.OrderStatus
	err := json.Unmarshal([]byte(request.Body), &orderStatus)

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf("Error: %s", err),
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
