#!/bin/bash

go test ./... -coverprofile=coverage.txt
go tool cover -html coverage.txt -o index.html

rm coverage.txt

xdg-open index.html