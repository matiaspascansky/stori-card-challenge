package transaction

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"stori-card-challenge/domain/transaction"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type TransactionRepository interface {
	GetTransactionsFromS3(bucket string, key string) ([]transaction.Transaction, error)
}

type transactionRepository struct {
	s3Client *s3.S3
}

func NewGetTransactionRepository(session *session.Session) *transactionRepository {
	return &transactionRepository{
		s3Client: s3.New(session),
	}
}

func (u *transactionRepository) GetTransactionsFromS3(bucket, key string) ([]transaction.Transaction, error) {
	// Perform S3 operation to get the object content
	getObjectOutput, err := u.s3Client.GetObjectWithContext(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer getObjectOutput.Body.Close()

	content, err := ioutil.ReadAll(getObjectOutput.Body)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(strings.NewReader(string(content)))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	transactions, err := validateAndProcessCSVRecords(records)

	if err != nil {
		return nil, errors.New("repository: error processing csv")
	}

	fmt.Printf("CSV Records: %+v\n", records)

	return transactions, nil
}

func validateAndProcessCSVRecords(records [][]string) ([]transaction.Transaction, error) {

	var transactions []transaction.Transaction

	for i, record := range records {

		if i == 0 {
			continue
		}

		if len(record) != 3 {
			return nil, errors.New("invalid number of fields in CSV record")
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, errors.New("invalid ID format")
		}
		//falta validar la fecha
		date := record[1]

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, errors.New("invalid Amount format")
		}

		tDto := TransactionDTO{
			ID:     id,
			Date:   date,
			Amount: amount,
		}

		t := FromDTOtoTransaction(tDto)

		transactions = append(transactions, t)

	}
	return transactions, nil

}
