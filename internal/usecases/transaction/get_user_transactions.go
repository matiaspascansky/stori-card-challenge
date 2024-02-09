package usecases

import (
	"errors"
	"stori-card-challenge/domain/transaction"
	transactionInfra "stori-card-challenge/internal/infrastructure/transaction"

	"github.com/aws/aws-sdk-go/aws/session"
)

type GetTransactionUsecase interface {
	GetTransactions(bucket, key string) ([]transaction.Transaction, error)
}

type getTransactionUsecase struct {
	transactionRepository transactionInfra.TransactionRepository
}

func NewGetTransactionUsecase(session *session.Session) *getTransactionUsecase {
	return &getTransactionUsecase{
		transactionRepository: transactionInfra.NewGetTransactionRepository(session),
	}
}

func (u *getTransactionUsecase) GetTransactions(bucket, key string) ([]transaction.Transaction, error) {
	// Get CSV content from S3
	transactions, err := u.transactionRepository.GetTransactionsFromS3(bucket, key)

	if err != nil {
		return nil, errors.New("usecase: error getting transactions from s3")

	}

	return transactions, nil

}
