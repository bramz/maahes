# Makefile for Maahes discord bot

depend:
	curl https://glide.sh/get | sh
	glide init ; glide update ; glide install			

build: | fmt test
	go build

test:
	go test ./...

fmt:
	go fmt ./...
