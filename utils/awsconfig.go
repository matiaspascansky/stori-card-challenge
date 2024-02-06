package utils

import (
	"encoding/json"
	"os"
)

type AWSConfig struct {
	AWSRegion    string `json:"aws_region"`
	AWSAccessKey string `json:"aws_access_key"`
	AWSSecretKey string `json:"aws_secret_key"`
	S3Bucket     string `json:"s3_bucket"`
	FilePath     string `json:"file_path"`
}

func ReadAWSConfig(filePath string) (AWSConfig, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return AWSConfig{}, err
	}
	defer file.Close()

	var config AWSConfig
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
