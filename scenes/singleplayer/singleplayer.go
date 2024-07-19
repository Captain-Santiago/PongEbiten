package singleplayer

import (
	"embed"
	"image/color"

	"github.com/Captain-Santiago/PongEbiten/scenes/player"
	"github.com/hajimehoshi/ebiten/v2"
)

type Singleplayer struct {
	assets *embed.FS

	player1                   *player.Player
	enemy                     *player.Player
	pointsPlayer, enemyPlayer int
}

func New(assets *embed.FS) *Singleplayer {
	return &Singleplayer{
		assets:       assets,
		player1:      player.NewPlayer(10*6, 80*6, 10*6, 35*6, color.RGBA{255, 0, 0, 255}),
		enemy:        player.NewPlayer(300*6, 80*6, 10*6, 35*6, color.RGBA{0, 0, 255, 255}),
		pointsPlayer: 0,
		enemyPlayer:  0,
	}
}

func (s *Singleplayer) Draw(screen *ebiten.Image) {
	// Draw Background
	screen.Fill(color.Black)

	// Draw Player 1
	s.player1.Draw(screen)

	// Draw Enemy
	s.enemy.Draw(screen)
}

func (s *Singleplayer) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.player1.PosY -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.player1.PosY += 10
	}

	return nil
}
