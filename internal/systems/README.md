# internal/systems

This package contains **pure transform functions** – stateless operations that
read and write component slices.

## Design principle

Each system is a function (or a struct with a single `Update` method) that
accepts component slices and a delta-time, and mutates them in place.  
No rendering, no I/O, no global state.

## Systems to implement

| File | System | Responsibility |
|------|--------|---------------|
| `physics.go` | `PhysicsSystem` | Integrate velocity → position each frame |
| `lifetime.go` | `LifetimeSystem` | Increment age; mark dead particles for removal |
| `emitter.go` | `EmitterSystem` | Spawn new particles from emitter configs |

## Execution order (in `Game.Update`)

1. `EmitterSystem` – spawn new particles
2. `PhysicsSystem` – move existing particles
3. `LifetimeSystem` – age & cull dead particles

## Rules

- Must not import `internal/ecs` (to avoid circular dependencies).
- Each system must be independently unit-testable.
