# Makefile for Maahes discord bot

depend:
	glide update ; glide install

build: | depend fmt test
	go build

test: | depend
	go test ./...

fmt: | depend
	go fmt ./...
