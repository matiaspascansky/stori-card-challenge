package usecases

import (
	"errors"
	"fmt"
	"stori-card-challenge/domain/transaction"
	transactionInfra "stori-card-challenge/internal/infrastructure/transaction"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
)

type ProcessTransactionsAndSendEmailUsecase interface {
	ProcessTransactionsAndSendEmail(transactions []transaction.Transaction) error
}

type processTransactionsAndSendEmailUsecase struct {
	emailSender transactionInfra.EmailSender
}

func NewProcessTransactionsAndSendEmailUsecase(session *session.Session) *processTransactionsAndSendEmailUsecase {
	return &processTransactionsAndSendEmailUsecase{
		emailSender: transactionInfra.NewGetEmailSender(session),
	}
}

func (u *processTransactionsAndSendEmailUsecase) ProcessTransactionsAndSendEmail(transactions []transaction.Transaction) error {

	status, err := processDataAndGenerateEmailContent(transactions)

	if err != nil {
		return errors.New("error processing data for email content creation")
	}

	fmt.Println("Total Balance:", status.TotalBalance)
	fmt.Println("Average Debit Amount:", status.AvgDebitAmount)
	fmt.Println("Average Credit Amount:", status.AvgCreditAmount)

	for year, months := range status.TransactionsGrouped.YearMonths {
		for month, transactions := range months {
			fmt.Printf("Year: %s, Month: %s, Transaction Count: %d\n", year, month, transactions)
		}
	}

	u.emailSender.SendEmail("asdasd")

	return nil

}

func processDataAndGenerateEmailContent(transactions []transaction.Transaction) (*transactionInfra.TransactionsStatus, error) {
	var totalBalance float64
	totalDebitTransactions := 0
	totalCreditTransaction := 0
	var sumDebitTransaction float64
	var sumCreditTransaction float64
	transactionCount := make(map[string]map[string]int)

	if len(transactions) == 0 {
		return nil, errors.New("process data: no transactions found")
	}

	for _, t := range transactions {
		totalBalance += t.Amount

		if t.Amount > 0 {
			totalCreditTransaction++
			sumCreditTransaction += t.Amount
		}
		if t.Amount < 0 {
			totalDebitTransactions++
			sumDebitTransaction += t.Amount
		}

		month := strings.Split(t.Date, "/")[0]
		year := strings.Split(t.Date, "/")[2]

		if transactionCount[year] == nil {
			transactionCount[year] = make(map[string]int)
		}

		transactionCount[year][month]++
	}

	tg := transactionInfra.TransactionsGroupedByYearAndMonth{
		YearMonths: transactionCount,
	}

	return &transactionInfra.TransactionsStatus{
		TotalBalance:        float64(totalBalance),
		AvgDebitAmount:      float64(sumDebitTransaction) / float64(totalDebitTransactions),
		AvgCreditAmount:     float64(sumCreditTransaction) / float64(totalCreditTransaction),
		TransactionsGrouped: tg,
	}, nil

}
