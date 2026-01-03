.PHONY: all deps generate benchmark report clean quick-benchmark verify

# Default target: run full benchmark and generate report.
all: generate benchmark report

# Install dependencies.
deps:
	go get github.com/goccy/go-json
	go get github.com/json-iterator/go
	go get github.com/mailru/easyjson
	go get github.com/tidwall/pretty
	go install github.com/mailru/easyjson/easyjson@latest

# Generate easyjson code.
generate:
	easyjson -all internal/testdata/types.go

# Run benchmarks with multiple iterations for accuracy.
benchmark: generate
	go test -bench=. -benchmem -count=5 ./benchmark/... 2>&1 | tee benchmark_results.txt

# Generate README report from benchmark results.
report:
	go run ./cmd/report/main.go benchmark_results.txt > README.md
	@echo "README.md generated successfully."

# Quick benchmark (single run, no report).
quick-benchmark:
	go test -bench=. -benchmem ./benchmark/...

# Verify that generated code is up to date.
verify: generate
	@git diff --exit-code internal/testdata/types_easyjson.go || \
		(echo "Generated files are out of date. Run 'make generate' and commit." && exit 1)

# Clean generated files.
clean:
	rm -f internal/testdata/types_easyjson.go
	rm -f benchmark_results.txt

# Build the report tool.
build-report:
	go build -o bin/report ./cmd/report

# Run tests to verify benchmark code compiles.
test:
	go test -v -run=^$$ -bench=^$$ ./benchmark/...
