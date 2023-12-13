# Build the program executable
.PHONY: all
all:
	go build -o term-ex

# Remove build files
.PHONY: clean
clean:
	rm -f term-ex

# Run tests
.PHONY: test
test:
	go test ./...

# Lint source code
.PHONY: lint
lint:
	golangci-lint run

# Format source code
.PHONY: format
format:
	go mod tidy
	go fmt
