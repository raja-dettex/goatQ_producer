run: build
	@./bin/goatQ_producer
build:
	@go build -o ./bin/goatQ_producer