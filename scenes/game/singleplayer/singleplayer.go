package singleplayer

import (
	"image/color"

	"github.com/Captain-Santiago/PongEbiten/scenes/game/player"
	"github.com/hajimehoshi/ebiten/v2"
)

type Singleplayer struct {
	player1                   *player.Player
	enemy                     *player.Player
	pointsPlayer, enemyPlayer int
}

func New() *Singleplayer {
	return &Singleplayer{
		player1:      player.NewPlayer(10, 80, 10, 35, color.RGBA{255, 0, 0, 255}),
		enemy:        player.NewPlayer(300, 80, 10, 35, color.RGBA{0, 0, 255, 255}),
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

func Move() {

}
