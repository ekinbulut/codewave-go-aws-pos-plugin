package main

import (
	"encoding/json"
	"order/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

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

	remoteId := request.PathParameters["remoteId"]
	body := request.Body

	var order models.Order

	err := json.Unmarshal([]byte(body), &order)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	}

	// TODO : bussiness logic

	// create a response
	resp := models.Response{
		RemoteResponse: models.RemoteResponse{
			RemoteOrderID: remoteId,
		},
	}

	// marshal response
	respBody, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(respBody),
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(handler)
}
