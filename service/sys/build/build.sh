#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o back_sys ../cmd/api/sys.go