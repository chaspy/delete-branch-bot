#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o delete-branch
zip handler.zip ./delete-branch
