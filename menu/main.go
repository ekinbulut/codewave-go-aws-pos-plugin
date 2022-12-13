package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Docs: https://integration-middleware.stg.restaurant-partners.com/apidocs/pos-plugin-api#tag/Plugin-Endpoints/operation/Trigger%20Menu%20Import

/*
Handler should return http status:
200 - OK
401 - Unauthorized
404 - Not Found
500 - Internal Server Error
502 - Bad Gateway
*/

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(handler)
}
