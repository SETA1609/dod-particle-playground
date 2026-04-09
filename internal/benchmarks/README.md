# internal/benchmarks

This package contains **micro-benchmarks** that measure the performance of the
DOD (Structure-of-Arrays) approach against a traditional
Array-of-Structures (AoS) layout.

## Files to implement

| File | Benchmark |
|------|-----------|
| `soa_bench_test.go` | SoA update loop – position + lifetime for N particles |
| `aos_bench_test.go` | AoS equivalent for comparison |
| `spawn_bench_test.go` | Cost of spawning / culling 10,000 particles |

## Running benchmarks

```bash
go test ./internal/benchmarks/... -bench=. -benchmem
# or via Makefile:
make bench
```

## Expected results

SoA should outperform AoS for large N due to better cache utilisation.
Document your findings in `README.md` once you have numbers.

## Rules

- Benchmark files end in `_test.go` and use `testing.B`.
- Do not import `cmd/` packages.
- Keep benchmarks deterministic (fixed random seed).
