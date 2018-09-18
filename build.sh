#!/bin/bash
rm -f ./delete-branch
rm -f handler.zip
GOOS=linux GOARCH=amd64 go build -o delete-branch
zip handler.zip ./delete-branch
