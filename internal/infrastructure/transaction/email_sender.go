package transaction

import (
	"bytes"
	"fmt"
	"html/template"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	recipient     = "matias.pascansky@gmail.com"
	emailTemplate = `
	<html>
	<head>
		<title>Transactions Status</title>
		<style>
			table {
				border-collapse: collapse;
				width: 100%;
			}
			th, td {
				border: 1px solid #dddddd;
				text-align: left;
				padding: 8px;
			}
			th {
				background-color: #f2f2f2;
			}
			.stori-card-logo {
				width: 400px;
				height: auto;
			}
		</style>
	</head>
	<body>
		<img src="https://s4-recruiting.cdn.greenhouse.io/external_greenhouse_job_boards/logos/400/955/600/original/logo_stori_1_(1).png?1690307995" alt="Story card Logo" class="stori-card-logo">
		<h1>Dear Customer</h1>
		<p>In this e-mail you will find information of your account status, if you have any concerns just contact us +55 7822 6646.</p>
		<h1>Transactions Status</h1>
		<p>Total Balance: ${{printf "%.2f" .TotalBalance}}</p>
		<p>Average Debit Amount: ${{printf "%.2f" .AvgDebitAmount}}</p>
		<p>Average Credit Amount: ${{printf "%.2f" .AvgCreditAmount}}</p>
		<h2>Transactions History:</h2>
		<table>
			<tr>
				<th>Year</th>
				<th>Month</th>
				<th>Transaction Count</th>
			</tr>
			{{range $year, $months := .TransactionsGrouped.YearMonths}}
				{{range $month, $count := $months}}
					<tr>
						<td>{{$year}}</td>
						<td>{{$month}}</td>
						<td>{{$count}}</td>
					</tr>
				{{end}}
			{{end}}
		</table>
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
	SendEmail(status *TransactionsStatus) error
}

type emailSender struct {
	sesClient *ses.SES
}

func NewGetEmailSender(session *session.Session) *emailSender {
	return &emailSender{
		sesClient: ses.New(session),
	}
}

func (e *emailSender) SendEmail(status *TransactionsStatus) error {

	emailContent, err := generateEmailContent(emailTemplate, status)
	if err != nil {
		log.Fatal("Error generating email content:", err)
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(recipient)},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(emailContent),
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

func generateEmailContent(templateStr string, data *TransactionsStatus) (string, error) {
	tmpl, err := template.New("email").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
