# internal/components

This package contains **pure data** – no logic, no methods (except maybe trivial constructors).

## Design principle (Structure-of-Arrays)

Each file defines **one component type** as a plain Go struct.  
Components are stored in separate slices in the ECS world, *not* bundled per entity, which keeps memory cache-friendly.

## Files to create

| File | Purpose |
|------|---------|
| `position.go` | `X, Y float32` world position |
| `velocity.go` | `VX, VY float32` pixels per second |
| `lifetime.go` | `Age, MaxAge float32` seconds |
| `color.go` | `R, G, B, A float32` normalised 0–1 |

## Rules

- No imports from `internal/systems` or `internal/ecs`.
- Structs should be kept as small as possible (prefer `float32` over `float64`).
