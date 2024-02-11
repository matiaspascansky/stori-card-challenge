package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"stori-card-challenge/domain/transaction"
	"stori-card-challenge/internal/infrastructure/topic"
	usecases "stori-card-challenge/internal/usecases/transaction"
	"stori-card-challenge/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

const (
	aws_config_path = "/var/task/aws_config.json"
)

type RequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// HandleAPIGatewayProxyRequest is the Lambda handler function.
func HandleAPIGatewayProxyRequest(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Read AWS configuration from JSON file
	config, err := utils.ReadAWSConfig(aws_config_path)
	if err != nil {
		log.Println("Error reading AWS config:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}

	// Create an AWS session
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(os.Getenv("aws_access_key"), os.Getenv("aws_secret_key"), ""),
	})
	if err != nil {
		log.Printf("Error creating session: %s", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Could not create aws session!",
		}, nil
	}

	var requestBody RequestBody
	if err := json.Unmarshal([]byte(r.Body), &requestBody); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Wrong type of request",
		}, nil
	}

	getTransactionsUsecase := usecases.NewGetTransactionUsecase(session)

	transactions, err := getTransactionsUsecase.GetTransactions(config.S3Bucket, config.ObjectKey)

	if err != nil {
		fmt.Println("Could not retrieve transactions:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "could not retrieve transactions!",
		}, nil
	}

	processAndSendEmailUsecase := usecases.NewProcessTransactionsAndSendEmailUsecase(session)

	transactionInfo, err := processAndSendEmailUsecase.ProcessTransactionsAndSendEmail(transactions, requestBody.Email)

	if err != nil {
		log.Println("error processing and sending email:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "broken!",
		}, nil
	}

	msgData := ToMsgData(requestBody, *transactionInfo)

	snsSender := topic.NewSnsSender(session, config.TopicArn)

	sendSnsMessageUsecase := usecases.NewSendMessageUsecase(snsSender)

	err = sendSnsMessageUsecase.Execute(msgData)

	if err != nil {
		log.Print("Warning: could not send message", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Email has been sent to user!",
	}, nil
}

func ToMsgData(r RequestBody, tInfo transaction.TransactionInformation) usecases.MsgData {
	return usecases.MsgData{
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		Email:        r.Email,
		Status:       tInfo.Status,
		TotalBalance: tInfo.TotalBalance,
	}
}
