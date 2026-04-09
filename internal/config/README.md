# internal/config

This package is responsible for **loading, validating, and hot-reloading** JSON
configuration files from `internal/data/`.

## Files to implement

| File | Purpose |
|------|---------|
| `types.go` | Go structs that mirror the JSON schema (`ParticleType`, `EmitterConfig`, …) |
| `loader.go` | `LoadParticleTypes`, `LoadEmitters`, `LoadScenes` functions |
| `hotreload.go` | Watch file for changes; reload on F5 key press |

## Key functions

```go
func LoadParticleTypes(path string) ([]ParticleType, error)
func LoadEmitters(path string) ([]EmitterConfig, error)
func (c *Config) Validate() error
func (c *Config) Reload() error  // called on F5
```

## Hot-reload behaviour

1. User presses **F5** in the game window.
2. `Config.Reload()` re-reads all JSON files from disk.
3. The new config is validated; if invalid the old config remains active.
4. The `ParticleSystem` picks up the new types on the next `Spawn` call.

## Rules

- Must not import `internal/ecs` or `internal/systems`.
- All file paths should be relative to the project root (pass via flag or embed).
