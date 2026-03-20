CONSUMERS=1

.PHONY: run vendor

vendor:
	go mod tidy
	go mod vendor


run:
	@go run . -consumers=$(CONSUMERS)