package usecases

import (
	"context"
	"log"
)

type SaveCsvUsecase interface {
	Execute(ctx context.Context, csvPath string) error
}

type saveCsvUsecase struct {
	//repositorioS3
}

func NewSaveCsvUsecase( /*repos3*/ ) *saveCsvUsecase {
	return &saveCsvUsecase{
		//s3repo
	}
}

func (s *saveCsvUsecase) Execute(ctx context.Context, csvPath string) error {

	log.Print("saving transactions.csv to s3 bucket")
	return nil
}
