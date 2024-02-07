package usecases

import (
	"stori-card-challenge/domain/transaction"
	transactionInfra "stori-card-challenge/internal/infrastructure/transaction"

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

	u.emailSender.SendEmail("asdasd")

	return nil

}
