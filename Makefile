.PHONY: clean-prepare
clean-prepare:
	@echo "Running clean-prepare"
	@echo "Run go mod tidy"
	@go mod tidy
	@echo "Run cp .env.example .env"
	@cp .env.example .env
	@echo "Finished"

.PHONY: test
test:
	@echo "Running test"
	@echo "Run go test -v ./..."
	@go test -v ./...

.PHONY: fmt
fmt:
	@echo "Running fmt"
	@echo "Run go fmt ./..."
	@go fmt ./...