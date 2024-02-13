# Stori card Challenge

This stori-card-challenge lambda where a Csv is got from S3 and the transactions are processed.

## Installation

To install this project in aws environment you need:

1) S3 configuration transactions-bucket-stori-card-challenge (bucket)
2) Create ECR repository called "docker-images" to upload images to run in lambda
3) Upload Csv, there is an example in /resources directory
4) Create Lambda in aws console configured with Docker image we are going to upload next
5) Push Dockerfile 
```bash 

## Build image

docker build -t stori-card-challenge-lambda:latest .

## Tag image

docker tag stori-card-challenge-lambda:latest 767397698520.dkr.ecr.us-east-1.amazonaws.com/docker-images

## Push Image to ECR service
docker push 767397698520.dkr.ecr.us-east-1.amazonaws.com/docker-images
```
5) Need to create SES identities for the emails in aws console (In the next version there is going to be a feature to create them automatically)

6) Create simple SNS Topic called "CreateUserAccountTopic" and in aws.json and replace Arn in aws_json


notes: 
- I am using us-east-1 region
- I also added the year to the csv to group not only by months, but by years as well.







