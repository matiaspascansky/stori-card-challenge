package usecases

import (
	"context"
	"errors"
	"stori-card-challenge/domain/transaction"
	"stori-card-challenge/internal/infrastructure/transaction/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProcessAccountTransactionsSendEmailTestSuite struct {
	suite.Suite
	ctx         context.Context
	emailSender *mocks.EmailSender
}

func (s *ProcessAccountTransactionsSendEmailTestSuite) SetupTest() {

	s.ctx = context.TODO()
	s.emailSender = new(mocks.EmailSender)

}

func (s *ProcessAccountTransactionsSendEmailTestSuite) TearDownTest() {
	s.emailSender.AssertExpectations(s.T())
}

func TestProcessAccountTransactionsSendEmail(t *testing.T) {
	suite.Run(t, new(ProcessAccountTransactionsSendEmailTestSuite))
}

func (s *ProcessAccountTransactionsSendEmailTestSuite) Test_Proccess_Transactions_And_Sends_Email() {
	s.T().Run("success", func(t *testing.T) {
		mockedTransactions := []transaction.Transaction{
			{
				ID:     1,
				Date:   "02/02/2012",
				Amount: 100.50,
			},
			{
				ID:     2,
				Date:   "02/02/2012",
				Amount: 50,
			}, {
				ID:     3,
				Date:   "02/02/2012",
				Amount: -10.50,
			},
		}
		s.emailSender.On("SendEmail", mock.Anything, mock.Anything).Return(nil)

		getTransactionUsecase := NewProcessTransactionsAndSendEmailUsecase(s.emailSender)

		inf, err := getTransactionUsecase.ProcessTransactionsAndSendEmail(mockedTransactions, "testEmail")

		assert.NoError(t, err)
		s.emailSender.AssertNumberOfCalls(s.T(), "SendEmail", 1)
		assert.Equal(t, float64(140), inf.TotalBalance)

	})

	s.TearDownTest()
	s.SetupTest()

	s.T().Run("process_transaction_error", func(t *testing.T) {
		mockedTransactions := []transaction.Transaction{}

		getTransactionUsecase := NewProcessTransactionsAndSendEmailUsecase(s.emailSender)

		_, err := getTransactionUsecase.ProcessTransactionsAndSendEmail(mockedTransactions, "testEmail")

		assert.ErrorContains(t, err, "process data: no transactions found")
		s.emailSender.AssertNumberOfCalls(s.T(), "SendEmail", 0)

	})

	s.TearDownTest()
	s.SetupTest()

	s.T().Run("emailNotSent", func(t *testing.T) {
		mockedTransactions := []transaction.Transaction{
			{
				ID:     1,
				Date:   "02/02/2012",
				Amount: 100.50,
			},
			{
				ID:     2,
				Date:   "02/02/2012",
				Amount: 50,
			}, {
				ID:     3,
				Date:   "02/02/2012",
				Amount: -10.50,
			},
		}
		s.emailSender.On("SendEmail", mock.Anything, mock.Anything).Return(errors.New("foo"))

		getTransactionUsecase := NewProcessTransactionsAndSendEmailUsecase(s.emailSender)

		_, err := getTransactionUsecase.ProcessTransactionsAndSendEmail(mockedTransactions, "testEmail")
		assert.ErrorContains(t, err, "error sending email to user")
		s.emailSender.AssertNumberOfCalls(s.T(), "SendEmail", 1)

	})

	s.TearDownTest()
	s.SetupTest()

}
