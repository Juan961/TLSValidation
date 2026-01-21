#!/bin/sh
export AWS_PROFILE=Personal

cd api/

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go

zip -r build.zip bootstrap templates/

cd ..

aws lambda update-function-code \
  --function-name GoValidateTLSAPI \
  --zip-file fileb://api/build.zip \
  --no-cli-pager > /dev/null && echo "✓ Lambda updated"
