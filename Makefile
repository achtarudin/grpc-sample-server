BIN_FILENAME=grpc-sample-server

.PHONY: install-deps
install-deps:
	go mod tidy

.PHONY: clean-bin
clean-bin:
	rm -rf ./bin

.PHONY: build
build: clean-bin install-deps
	go build -o ./bin/$(BIN_FILENAME) ./cmd/server

.PHONY: execute
execute: build
	./bin/$(BIN_FILENAME)

.PHONY: dev-server
dev-server:
	gow run ./cmd/server

.PHONY: dev-cli
dev-cli:
	gow run ./cmd/cli