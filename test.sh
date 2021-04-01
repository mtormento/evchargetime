#!/bin/bash

go test -timeout 30s -cover github.com/mtormento/evchargetime/go/calc
go test -timeout 30s -cover github.com/mtormento/evchargetime/go/fmt
