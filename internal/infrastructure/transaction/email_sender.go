package transaction

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	recipient = "matias.pascansky@gmail.com"
	body      = "mensajito body"
	subject   = "Stori Card Transactions Status"
	sender    = "matias.pascansky@gmail.com"
)

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
