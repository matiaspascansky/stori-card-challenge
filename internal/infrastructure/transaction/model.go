package transaction

import "stori-card-challenge/domain/transaction"

type TransactionDTO struct {
	ID     int     `json:"id"`
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}

func FromDTOtoTransaction(dto TransactionDTO) transaction.Transaction {
	return transaction.Transaction{
		ID:     dto.ID,
		Date:   dto.Date,
		Amount: dto.Amount,
	}
}

type TransactionsStatus struct {
	TotalBalance        float64                           `json:"total_balance"`
	AvgDebitAmount      float64                           `json:"avg_debit_amount"`
	AvgCreditAmount     float64                           `json:"avg_credit_amount"`
	TransactionsGrouped TransactionsGroupedByYearAndMonth `json:"transactions_grouped"`
	Status              string                            `json:"status"`
}

type TransactionsGroupedByYearAndMonth struct {
	YearMonths map[string]map[string]int
}
