# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /go/src/stori-card-challenge

# Copy the entire project (including source code and go.mod/go.sum) into the container
COPY . .

# Explicitly set the module path (replace 'example.com' with your actual module path)
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build the Go application for the Lambda execution environment
RUN go build -o main .

# Command to run the Lambda function
CMD [ "./main" ]
EXPOSE 8080
