#!/bin/bash

go run cmd/bun/main.go db-cart $1
go run cmd/bun/main.go db-invoice $1
go run cmd/bun/main.go db-order $1
go run cmd/bun/main.go db-payment $1
