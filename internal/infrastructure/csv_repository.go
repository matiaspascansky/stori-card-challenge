package infrastructure

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type CsvRepository interface {
	GetCsvFromS3(bucket string, key string) ([]byte, error)
}

type csvRepository struct {
	s3Client *s3.S3
}

func NewGetCsvRepository(session *session.Session) *csvRepository {
	return &csvRepository{
		s3Client: s3.New(session),
	}
}

func (u *csvRepository) GetCsvFromS3(bucket, key string) ([]byte, error) {
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

	if err := validateCSVRecords(records); err != nil {
		return nil, err
	}

	fmt.Printf("CSV Records: %+v\n", records)

	return nil, nil
}

func validateCSVRecords(records [][]string) error {
	for _, record := range records {
		if len(record) != 3 {
			return errors.New("invalid number of fields in CSV record")
		}

		_, err := strconv.Atoi(record[0])
		if err != nil {
			return errors.New("invalid ID format")
		}

		_, err = strconv.ParseFloat(record[2], 64)
		if err != nil {
			return errors.New("invalid Amount format")
		}

	}

	return nil
}
