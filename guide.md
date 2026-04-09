# DOD Particle Playground – Development Guide

## Phase 0: Project Setup from Absolute Zero (1 day)

### Issue 0.1 – Create the repository and Go module

```bash
mkdir dod-particle-playground && cd dod-particle-playground
git init
go mod init github.com/yourusername/dod-particle-playground  # replace with your GitHub name
```

### Issue 0.2 – Add Ebitengine dependency

```bash
go get github.com/hajimehoshi/ebiten/v2
```

Create the folder structure exactly as shown below (empty files for now):

```text
dod-particle-playground/
├── cmd/
│   └── playground/
│       └── main.go          ← entry point
├── internal/
│   ├── components/          ← SoA data only
│   ├── systems/             ← pure transform functions
│   ├── ecs/                 ← world + entity manager
│   ├── data/                ← JSON configs (will create later)
│   ├── config/              ← loaders + hot-reload
│   └── benchmarks/          ← later
├── pkg/
│   └── dod/                 ← reusable helpers
├── go.mod
├── go.sum
├── README.md
└── Makefile                 ← optional but useful
```

### Issue 0.3 – Create the minimal Ebitengine "Hello World" game loop

In `cmd/playground/main.go` create a struct `Game` that implements the three required Ebitengine methods (empty for now):

```go
type Game struct{}

func (g *Game) Update() error { return nil }
func (g *Game) Draw(screen *ebiten.Image) {}
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) { return 0, 0 }
```

Add the `main()` function that starts the game (look up the official Ebitengine "Hello World" example and copy only the boilerplate structure).

### Issue 0.4 – Run it and confirm a blank window opens

```bash
go run ./cmd/playground
```

You should see a window. Close it with ESC or window close button.

### Issue 0.5 – Add basic window settings (title, size 1280×720, resizable) and commit.

---

## Phase 1: DOD Foundation – Structure-of-Arrays Particle System (2–3 days)

### Issue 1.1 – Create the first components (empty files in `internal/components/`)

`position.go`, `velocity.go`, `lifetime.go`, `color.go`

### Issue 1.2 – Implement the DOD ParticleSystem (SoA style)

In `internal/ecs/particle_system.go` create a struct with separate slices (no code yet – just the struct and a `New` function).

### Issue 1.3 – Add Spawn and the Update loop skeleton

Create these empty methods:

```go
func (ps *ParticleSystem) Spawn(...) {}
func (ps *ParticleSystem) Update(dt float64) {}
```

Inside `Update` you will later do the position/velocity/lifetime transforms + swap-and-pop removal.

### Issue 1.4 – Connect ParticleSystem to the Ebitengine Game loop

Add the system to the `Game` struct and call `Update` and a future `Draw` every frame.

### Issue 1.5 – Implement a very simple Draw stub (just clear the screen to black for now).

### Issue 1.6 – Spawn 10,000 particles at startup and verify the window still runs at 60 FPS (use `ebiten.SetTPS(0)` and check FPS in title).

---

## Phase 2: Data-Driven Layer – JSON Configs (3–4 days)

### Issue 2.1 – Create the first JSON file

`internal/data/particles.json` (empty object for now)

### Issue 2.2 – Define `ParticleType` struct in `internal/config/types.go`

### Issue 2.3 – Write a Config loader

Empty functions:

```go
func LoadParticleTypes(path string) ([]ParticleType, error) { return nil, nil }
func (c *Config) Validate() error {}
```

### Issue 2.4 – Add hot-reload support (watch the JSON file and reload on F5 key press).

### Issue 2.5 – Make Spawn use data from JSON (different particle types with different colors, speeds, lifetimes).

### Issue 2.6 – Test: Change the JSON → press F5 → see new particle behavior without restarting the program.

---

## Phase 3: Systems & Full ECS Architecture (4–5 days)

### Issue 3.1 – Create separate System interfaces in `internal/systems/`

### Issue 3.2 – Split logic into these systems (each as empty func):

- `PhysicsSystem`
- `LifetimeSystem`
- `EmitterSystem`

### Issue 3.3 – Implement archetype storage (simple grouping of components).

### Issue 3.4 – Add multiple emitters (fire, explosion, fountain) loaded from `emitters.json`.

### Issue 3.5 – Connect all systems to the main `Game.Update` in the correct order.

---

## Phase 4: Polish & Advanced DOD (3–4 days)

### Issue 4.1 – Add spatial queries (e.g. particles near mouse cursor).

### Issue 4.2 – Implement benchmark suite in `internal/benchmarks/` (compare SoA vs hypothetical AoS).

### Issue 4.3 – Add simple mouse interaction (click to spawn explosion).

### Issue 4.4 – Nice visuals: trails, alpha blending, different blend modes.

### Issue 4.5 – Create 3 preset scenes in JSON (fireworks, rain, galaxy).

---

## Phase 5: Final Touches & Reflection (1–2 days)

### Issue 5.1 – Write a detailed README explaining what you learned about DOD + Data-Driven.

### Issue 5.2 – Add Makefile targets (`make run`, `make bench`, `make clean`).

### Issue 5.3 – (Optional) Push to GitHub and share the link if you want feedback.

---

## How to use this plan

- Do **one issue per session**.
- After each issue: `git add . && git commit -m "milestone X - issue Y: description"`
- When you get stuck on any issue, search for:
  - `"Ebitengine [method name] example"`
  - `"Go Structure of Arrays particle system"`
  - `"Go JSON config hot reload"`
