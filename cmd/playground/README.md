# cmd/playground

This is the **entry point** of the application.

## What belongs here

- `main.go` – initialises the Ebitengine window, wires up the top-level `Game`
  struct, and calls `ebiten.RunGame`.

## Key responsibilities

- Set window title, size (1280×720), and resizable flag.
- Instantiate the `ParticleSystem` (or ECS `World`) and hand it to `Game`.
- Keep this file as thin as possible – all logic lives in `internal/`.

## Getting started

```bash
go run ./cmd/playground
```
