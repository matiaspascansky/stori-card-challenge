package handler

import (
	"context"
	"fmt"
	usecases "stori-card-challenge/internal/usecases/csv"
	"stori-card-challenge/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	aws_config_path = "/var/task/aws_config.json"
)

// HandleAPIGatewayProxyRequest is the Lambda handler function.
func HandleAPIGatewayProxyRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Read AWS configuration from JSON file
	config, err := utils.ReadAWSConfig(aws_config_path)
	if err != nil {
		fmt.Println("Error reading AWS config:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}

	// Create an AWS session
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(config.AWSAccessKey, config.AWSSecretKey, ""),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}
	getTransactionsUsecase := usecases.NewGetTransactionUsecase(session)

	transactions, err := getTransactionsUsecase.GetTransactions(config.S3Bucket, config.ObjectKey)
	fmt.Println("handler: OK TRANSACTIONS", err)

	processAndSendEmailUsecase := usecases.NewProcessTransactionsAndSendEmailUsecase(session)

	err = processAndSendEmailUsecase.ProcessTransactionsAndSendEmail(transactions)

	if err != nil {
		fmt.Println("error processing and sending email:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}

	fmt.Print(transactions)

	if err != nil {
		fmt.Println("handler: Error in csv usecase", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Csv is processed completely!",
	}, nil
}
