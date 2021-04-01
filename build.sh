#!/bin/sh

CGO_ENABLED=0 GOOS=linux go build -o bin/evchargetime -a -installsuffix cgo -ldflags '-extldflags "-static"' go/charge_time.go
