package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) UpdatePlayerMovement() {
	// Move Player 1

	// Reset movement
	g.SceneManager.MoveX1 = 0
	g.SceneManager.MoveY1 = 0

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.SceneManager.MoveY1 += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.SceneManager.MoveY1 -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.SceneManager.MoveX1 += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.SceneManager.MoveX1 -= 1
	}

	// Move Player 2 (if applicable)

	// Reset movement
	g.SceneManager.MoveX2 = 0
	g.SceneManager.MoveY2 = 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.SceneManager.MoveY2 += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.SceneManager.MoveY2 -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.SceneManager.MoveX2 += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.SceneManager.MoveX2 -= 1
	}
}
