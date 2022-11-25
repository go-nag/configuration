.PHONY: clean-prepare
clean-prepare:
	@go mod tidy

.PHONY: test
test:
	@go test -v ./...

.PHONY: test-cov
test-cov:
	@go test -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: fmt
fmt:
	@go fmt ./...