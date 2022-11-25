.PHONY: clean-prepare
clean-prepare:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./...

.PHONY: fmt
fmt:
	@go fmt ./...