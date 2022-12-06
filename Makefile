.PHONY: build
build:
		go build -v ./cmd/server

.PHONY: run-server
run-server:
		go run -v ./cmd/server/main.go


##.PHONY: run-client
##run-client:
##		go run -v ./cmd/client/main.go


.DEFAULT_GOAL := build