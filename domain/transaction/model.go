package transaction

type Transaction struct {
	ID     int     `json:"Id"`
	Date   string  `json:"Date"`
	Amount float64 `json:"Amount"`
}

type TransactionInformation struct {
	TotalBalance float64 `json:"total_balance"`
	Status       string  `json:"status"`
}
