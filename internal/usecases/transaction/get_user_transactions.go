package usecases

import (
	"errors"
	"stori-card-challenge/domain/transaction"
	transactionInfra "stori-card-challenge/internal/infrastructure/transaction"

	"github.com/aws/aws-sdk-go/aws/session"
)

type TransactionUsecase interface {
	GetTransactions(bucket, key string) ([]transaction.Transaction, error)
}

type transactionUsecase struct {
	transactionRepository transactionInfra.TransactionRepository
}

func NewGetTransactionUsecase(session *session.Session) *transactionUsecase {
	return &transactionUsecase{
		transactionRepository: transactionInfra.NewGetTransactionRepository(session),
	}
}

func (u *transactionUsecase) GetTransactions(bucket, key string) ([]transaction.Transaction, error) {
	// Get CSV content from S3
	transactions, err := u.transactionRepository.GetTransactionsFromS3(bucket, key)

	if err != nil {
		return nil, errors.New("usecase: error getting transactions from s3")

	}

	return transactions, nil

}
