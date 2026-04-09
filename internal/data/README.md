# internal/data

This directory stores **static JSON configuration files** that define particle
types, emitter presets, and scene layouts.

## Files to create

| File | Purpose |
|------|---------|
| `particles.json` | Library of named particle types (color, speed, lifetime, …) |
| `emitters.json` | Named emitter configs (fire, explosion, fountain, …) |
| `scenes.json` | Three preset scenes: fireworks, rain, galaxy |

## Example `particles.json` shape

```json
{
  "types": [
    {
      "name": "spark",
      "color": [1.0, 0.8, 0.2, 1.0],
      "speedMin": 80,
      "speedMax": 200,
      "lifetimeMin": 0.5,
      "lifetimeMax": 1.5
    }
  ]
}
```

## Rules

- JSON files are **read-only at runtime** (loaded by `internal/config`).
- Keep keys in `camelCase` to match Go struct field tags.
- Validate files with `go run ./cmd/playground -validate` once that flag exists.
