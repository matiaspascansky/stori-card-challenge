package usecases

import (
	"errors"
	"stori-card-challenge/domain/transaction"
	transactionInfra "stori-card-challenge/internal/infrastructure/transaction"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
)

type ProcessTransactionsAndSendEmailUsecase interface {
	ProcessTransactionsAndSendEmail(transactions []transaction.Transaction, email string) error
}

type processTransactionsAndSendEmailUsecase struct {
	emailSender transactionInfra.EmailSender
}

func NewProcessTransactionsAndSendEmailUsecase(session *session.Session) *processTransactionsAndSendEmailUsecase {
	return &processTransactionsAndSendEmailUsecase{
		emailSender: transactionInfra.NewGetEmailSender(session),
	}
}

func (u *processTransactionsAndSendEmailUsecase) ProcessTransactionsAndSendEmail(transactions []transaction.Transaction, email string) (*transaction.TransactionInformation, error) {

	tStatus, err := processDataAndCalculateStatus(transactions)

	tInfo := &transaction.TransactionInformation{
		TotalBalance: tStatus.TotalBalance,
		Status:       tStatus.Status,
	}

	if err != nil {
		return nil, errors.New("error processing data for email content creation")
	}

	u.emailSender.SendEmail(tStatus, email)

	return tInfo, err

}

func processDataAndCalculateStatus(transactions []transaction.Transaction) (*transactionInfra.TransactionsStatus, error) {
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
		if t.Amount == 0 {
			continue
		}
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

	status := ""
	if totalBalance < 0 {
		status = "in debt"
	} else {
		status = "debt-free"
	}

	return &transactionInfra.TransactionsStatus{
		TotalBalance:        float64(totalBalance),
		AvgDebitAmount:      float64(sumDebitTransaction) / float64(totalDebitTransactions),
		AvgCreditAmount:     float64(sumCreditTransaction) / float64(totalCreditTransaction),
		TransactionsGrouped: tg,
		Status:              status,
	}, nil

}
