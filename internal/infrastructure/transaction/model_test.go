package transaction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromDTO(t *testing.T) {
	tDto := TransactionDTO{
		ID:     1,
		Date:   "02/03/2024",
		Amount: 35.3,
	}

	transaction := FromDTOtoTransaction(tDto)

	assert.Equal(t, tDto.ID, transaction.ID)
	assert.Equal(t, tDto.Date, transaction.Date)
	assert.Equal(t, tDto.Amount, transaction.Amount)

}
