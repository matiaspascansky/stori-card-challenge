package handler

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

// HandleAPIGatewayProxyRequest is the Lambda handler function.
func HandleAPIGatewayProxyRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, Lambda!",
	}, nil
}

// HealthCheckHandler is the health check handler function.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
