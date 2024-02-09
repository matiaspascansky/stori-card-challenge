package main

import (
	"stori-card-challenge/cmd/web/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Register Lambda handlers
	lambda.Start(handler.HandleAPIGatewayProxyRequest)

}
