BINARY   := playground
CMD      := ./cmd/playground
BENCH    := ./internal/benchmarks/...

.PHONY: run bench clean

## run: build and run the application
run:
	go run $(CMD)

## bench: run all benchmarks with memory stats
bench:
	go test $(BENCH) -bench=. -benchmem

## clean: remove build artifacts
clean:
	go clean -cache
	rm -f $(BINARY)
