package usecases

import (
	"stori-card-challenge/internal/infrastructure/topic"

	"github.com/pkg/errors"
)

type SendMessageUsecase interface {
	Execute(MsgData MsgData) error
}

type sendMessageUsecase struct {
	snsSender topic.SnsSender
}

func NewSendMessageUsecase(snsSender topic.SnsSender) *sendMessageUsecase {
	return &sendMessageUsecase{
		snsSender: snsSender,
	}
}

type MsgData struct {
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	TotalBalance float64 `json:"total_balance"`
	Status       string  `json:"status"`
	Email        string  `json:"email"`
}

func (m *MsgData) ToTopicMessage() topic.TopicMessage {
	return topic.TopicMessage{
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		TotalBalance: m.TotalBalance,
		Status:       m.Status,
	}
}

func (s *sendMessageUsecase) Execute(msgData MsgData) error {
	tm := msgData.ToTopicMessage()
	err := s.snsSender.Execute(tm)

	if err != nil {
		return errors.Wrap(err, "usecase: cannot send msg to sns")
	}
	return err
}
