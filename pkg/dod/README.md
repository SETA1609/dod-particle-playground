# pkg/dod

This package contains **reusable, dependency-free helpers** related to
Data-Oriented Design patterns that could be used in other projects.

## Planned utilities

| File | Purpose |
|------|---------|
| `pool.go` | Generic fixed-capacity pool with swap-and-pop removal |
| `math.go` | Fast `float32` math helpers (lerp, clamp, random range) |
| `rand.go` | Lightweight, seedable random number generator for particle spread |

## Design goals

- Zero dependencies outside the Go standard library.
- All types and functions generic where beneficial (`pool.go` uses `[T any]`).
- 100% unit-tested.

## Example usage

```go
import "github.com/yourusername/dod-particle-playground/pkg/dod"

pool := dod.NewPool[MyComponent](10_000)
pool.Add(item)
pool.RemoveAt(i)  // swap-and-pop
```

## Rules

- Must not import anything from `internal/`.
- Exported API should be stable once defined.
