package usecases

import (
	"context"
	"stori-card-challenge/domain/transaction"
	"stori-card-challenge/internal/infrastructure/transaction/mocks"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GetUserTransactionsTestSuite struct {
	suite.Suite
	ctx              context.Context
	transactionsRepo *mocks.TransactionRepository
	session          *session.Session
}

func (s *GetUserTransactionsTestSuite) SetupTest() {

	awsConfig := aws.Config{
		Region:      aws.String("test-region"),
		Credentials: credentials.NewStaticCredentials("testKey", "testSecret", ""),
	}

	s.ctx = context.TODO()
	s.transactionsRepo = new(mocks.TransactionRepository)
	s.session, _ = session.NewSession(&awsConfig)

}

func (s *GetUserTransactionsTestSuite) TearDownTest() {
	s.transactionsRepo.AssertExpectations(s.T())
}

func TestGetTransactions(t *testing.T) {
	suite.Run(t, new(GetUserTransactionsTestSuite))
}

func (s *GetUserTransactionsTestSuite) Test_GetTransactions() {
	s.T().Run("success", func(t *testing.T) {
		mockedTransactions := []transaction.Transaction{
			{
				ID:     1,
				Date:   "02/02/2012",
				Amount: 100.50,
			},
		}

		bucketName := "test"
		key := "testKey"

		s.transactionsRepo.On("GetTransactionsFromS3", mock.Anything, mock.Anything).Return(mockedTransactions, nil)

		getTransactionUsecase := NewGetTransactionUsecase(s.transactionsRepo)

		trans, err := getTransactionUsecase.GetTransactions(bucketName, key)

		s.transactionsRepo.AssertNumberOfCalls(s.T(), "GetTransactionsFromS3", 1)

		assert.NoError(t, err)
		assert.Equal(t, mockedTransactions[0].Amount, trans[0].Amount)

	})
}
