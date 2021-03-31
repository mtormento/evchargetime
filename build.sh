#!/bin/sh

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' go/charge_time.go
mv charge_time bin
