.PHONY: clean-prepare
clean-prepare:
	@go mod tidy
	@cp .env.example .env

.PHONY: test
test:
	@go test -v ./...

.PHONY: fmt
fmt:
	@go fmt ./...