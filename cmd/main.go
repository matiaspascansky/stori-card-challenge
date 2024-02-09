package main

import (
	"log"
	"stori-card-challenge/cmd/web/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Print("este es el punto de entrada")
	// Register Lambda handlers
	lambda.Start(handler.HandleAPIGatewayProxyRequest)

}
