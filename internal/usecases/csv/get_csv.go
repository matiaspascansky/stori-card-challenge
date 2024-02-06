package usecase

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
)

type CsvUsecase interface {
	ProcessCsvFromS3(bucket, key string) error
}

type csvUsecase struct {
	csvRepository CsvRepository
}

func NewCsvUsecase(session *session.Session) CsvUsecase {
	return &csvUsecase{
		csvRepository: NewGetCsvRepository(session),
	}
}

func (u *csvUsecase) ProcessCsvFromS3(bucket, key string) error {
	// Get CSV content from S3
	content, err := u.csvRepository.GetCsvFromS3(bucket, key)
	if err != nil {
		return err
	}

	// Process the CSV content
	if err := processCSV(content); err != nil {
		return err
	}

	return nil
}

func processCSV(content []byte) error {
	// Process the CSV content as needed
	fmt.Printf("CSV Content: %s\n", content)
	return nil
}
