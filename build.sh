#!/bin/bash
rm -f ./delete-branch
rm -f handler.zip
GOOS=linux GOARCH=amd64 go build -o delete-branch
zip handler.zip ./delete-branch ./private-key.pem
aws lambda update-function-code \
  --function-name delete_branch_bot \
  --zip-file fileb://handler.zip \
  --region ap-northeast-1
