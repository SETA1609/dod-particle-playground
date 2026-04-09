package main

// cmd/playground/main.go
//
// Entry point for the DOD Particle Playground.
// Phase 0 stub – a blank Ebitengine window (1280×720, resizable).
//
// Run:
//   go run ./cmd/playground

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements the ebiten.Game interface.
type Game struct{}

// Update is called every tick (default 60 TPS).
func (g *Game) Update() error {
	return nil
}

// Draw is called every frame.
func (g *Game) Draw(screen *ebiten.Image) {}

// Layout returns the logical screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("DOD Particle Playground")
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
