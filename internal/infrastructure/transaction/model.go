package transaction

import "stori-card-challenge/domain/transaction"

type TransactionDTO struct {
	ID     int     `json:"Id"`
	Date   string  `json:"Date"`
	Amount float64 `json:"Amount"`
}

func FromDTOtoTransaction(dto TransactionDTO) transaction.Transaction {
	return transaction.Transaction{
		ID:     dto.ID,
		Date:   dto.Date,
		Amount: dto.Amount,
	}
}
