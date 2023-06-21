-include .env

all: proto build cover gen
.PHONY: all

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/proto/product.proto

build:
	go build product/cmd/main

cover:
	go test -short -count=1 -race -coverprofile=coverage.out product/internal/server
	go tool cover -html=coverage.out

gen:
	mockgen -source=internal/repository/repository.go \
	-destination=internal/repository/mocks/mock_repository.go
