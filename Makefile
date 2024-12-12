include .env
.PHONY: m-generate

dev:
	go build -o ./tmp/main cmd/server/main.go; ./tmp/main

wire:
	cd internal/injector && wire && cd ../..
	