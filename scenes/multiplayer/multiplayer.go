package multiplayer

import (
	"embed"
	"image/color"

	"github.com/Captain-Santiago/PongEbiten/scenes/player"
	"github.com/hajimehoshi/ebiten/v2"
)

type Multiplayer struct {
	assets *embed.FS

	player1 *player.Player
	player2 *player.Player
}

func New(assets *embed.FS) *Multiplayer {
	return &Multiplayer{
		assets:  assets,
		player1: player.NewPlayer(10*6, 80*6, 10*6, 35*6, color.RGBA{255, 0, 0, 255}),
		player2: player.NewPlayer(300*6, 80*6, 10*6, 35*6, color.RGBA{0, 0, 255, 255}),
	}
}

func (m *Multiplayer) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		m.player1.PosY -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		m.player1.PosY += 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		m.player2.PosY -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		m.player2.PosY += 10
	}

	return nil
}

func (m *Multiplayer) Draw(screen *ebiten.Image) {
	// Draw Background
	screen.Fill(color.Black)

	// Draw Player 1
	m.player1.Draw(screen)

	// Draw player 2
	m.player2.Draw(screen)
}
