#!/bin/bash

# Set target for AWS Lambda (Linux x86_64)
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

# Remove previous zip (if any)
rm -f function.zip

# Zip the binary with the name expected by Lambda
zip function.zip bootstrap