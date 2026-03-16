# Variables
BINARY_DIR=bin
BENCHMARK_BIN=$(BINARY_DIR)/goliath-benchmark

# Build the benchmark binary
build:
	@echo "Building goliath benchmark..."
	@mkdir -p $(BINARY_DIR)
	@go build -o $(BENCHMARK_BIN) ./cmd/benchmark/main.go
	@echo "Build complete: $(BENCHMARK_BIN)"

# Run the compiled binary
run: build
	@echo "Running goliath benchmark..."
	@./$(BENCHMARK_BIN)

# Run tests with GC tracing
test:
	@echo "Running tests with GC tracing..."
	@GODEBUG=gctrace=1 go test -bench=. ./pkg/memory/

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BINARY_DIR)