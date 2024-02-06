package usecases

import (
	"errors"
	"fmt"
	"stori-card-challenge/internal/infrastructure/transaction"

	"github.com/aws/aws-sdk-go/aws/session"
)

type TransactionUsecase interface {
	ProcessTransactions(bucket, key string) error
}

type transactionUsecase struct {
	transactionRepository transaction.TransactionRepository
}

func NewGetTransactionUsecase(session *session.Session) *transactionUsecase {
	return &transactionUsecase{
		transactionRepository: transaction.NewGetTransactionRepository(session),
	}
}

func (u *transactionUsecase) ProcessTransactions(bucket, key string) error {
	// Get CSV content from S3
	transactions, err := u.transactionRepository.GetTransactionsFromS3(bucket, key)
	for _, transaction := range transactions {
		fmt.Printf("ID: %d, Date: %s, Amount: %.2f\n", transaction.ID, transaction.Date, transaction.Amount)
	}
	if err != nil {
		errors.New("usecase: error getting transactions from s3")
	}

	return nil

}

type CsvResponse struct {
}
