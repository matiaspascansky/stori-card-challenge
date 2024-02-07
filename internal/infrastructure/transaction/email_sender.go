package transaction

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	recipient = "matias.pascansky@gmail.com"
	body      = `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset="UTF-8">
		<title>Email Template</title>
	</head>
	<body>
		<h1>Hello, {{.Name}}!</h1>
		<p>{{.Message}}</p>
	</body>
	</html>
	`
	subject = "Stori Card Transactions Status"
	sender  = "matias.pascansky@gmail.com"
)

type EmailData struct {
	//todo lo que vamos a reemplazar en el mail
}

type EmailSender interface {
	SendEmail(content string) error
}

type emailSender struct {
	sesClient *ses.SES
}

func NewGetEmailSender(session *session.Session) *emailSender {
	return &emailSender{
		sesClient: ses.New(session),
	}
}

func (e *emailSender) SendEmail(content string) error {
	// Specify the email input
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	// Send the email
	result, err := e.sesClient.SendEmail(input)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	fmt.Println("Email sent successfully:", result)
	return nil
}
