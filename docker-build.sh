#!/bin/bash

go mod download
CGO_ENABLED=0 GOOS=linux go build -o http-mock-service
docker buildx build -t dominikcebula/http-mock-service:latest .
