package usecases

import (
	"errors"
	"stori-card-challenge/domain/transaction"
	transactionInfra "stori-card-challenge/internal/infrastructure/transaction"
)

type GetTransactionUsecase interface {
	GetTransactions(bucket, key string) ([]transaction.Transaction, error)
}

type getTransactionUsecase struct {
	transactionRepository transactionInfra.TransactionRepository
}

func NewGetTransactionUsecase(transactionRepository transactionInfra.TransactionRepository) *getTransactionUsecase {
	return &getTransactionUsecase{
		transactionRepository: transactionRepository,
	}
}

func (u *getTransactionUsecase) GetTransactions(bucket, key string) ([]transaction.Transaction, error) {

	transactions, err := u.transactionRepository.GetTransactionsFromS3(bucket, key)

	if err != nil {
		return nil, errors.New("usecase: error getting transactions from s3")

	}

	return transactions, nil

}
