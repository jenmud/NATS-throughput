CONSUMERS=1


.PHONY: run vendor build


vendor:
	go mod tidy
	go mod vendor


run:
	@go run . -consumers=$(CONSUMERS)


build:
	@go build -o builds/NATS-throughput