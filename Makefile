
.PHONY: test

test:
	go test -v -count 1 ./...

cover:
	go test -cover -coverprofile ./...
	go test -cover -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html

lint:
	golangci-lint run
