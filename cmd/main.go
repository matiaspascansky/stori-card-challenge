package main

import (
	"stori-card-challenge/internal/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Register Lambda handlers
	lambda.Start(handler.HandleAPIGatewayProxyRequest)

	// For local testing with a health check endpoint
	//http.HandleFunc("/health", handler.HealthCheckHandler)
	//fmt.Println("Health check server started on :8080")
	//http.ListenAndServe(":8080", nil)
}
