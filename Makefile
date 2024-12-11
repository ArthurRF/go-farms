include .env
.PHONY: m-generate

dev:
	make m-apply; go build -o ./tmp/main cmd/server/main.go; ./tmp/main

m-apply:
	atlas migrate hash; atlas migrate apply --url ${DB_URL}