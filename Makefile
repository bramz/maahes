# Makefile for Maahes discord bot

build: | fmt test
	go build

test:
	go test ./...

fmt:
	go fmt ./...