# internal/ecs

This package contains the **world and entity manager** – the glue that owns all
component slices and orchestrates systems.

## Key types to implement

| File | Type | Responsibility |
|------|------|---------------|
| `world.go` | `World` | Owns all component arrays; exposes `AddEntity`, `RemoveEntity` |
| `particle_system.go` | `ParticleSystem` | SoA particle pool: `Spawn`, `Update`, `Draw` |
| `entity.go` | `EntityID` | Lightweight integer handle |

## Structure-of-Arrays layout

```go
type ParticleSystem struct {
    X, Y     []float32  // positions
    VX, VY   []float32  // velocities
    Age      []float32
    MaxAge   []float32
    R, G, B, A []float32
    Count    int
}
```

## Swap-and-pop removal

When a particle dies, swap it with the last alive particle and decrement
`Count`.  This avoids gaps and keeps memory contiguous.

## Rules

- May import `internal/components` and `internal/systems`.
- Must not import `cmd/` packages.
