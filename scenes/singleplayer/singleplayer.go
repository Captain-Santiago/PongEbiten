package singleplayer

import (
	"embed"
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/Captain-Santiago/PongEbiten/scenes/ball"
	"github.com/Captain-Santiago/PongEbiten/scenes/enemy"
	"github.com/Captain-Santiago/PongEbiten/scenes/player"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Singleplayer struct {
	assets *embed.FS

	player1      *player.Player
	enemy        *enemy.Enemy
	ball         *ball.Ball
	pointsPlayer int
	enemyPlayer  int
	background   *ebiten.Image
}

func New(assets *embed.FS) *Singleplayer {
	var err error
	background, _, err := ebitenutil.NewImageFromFileSystem(assets, "assets/game/ceu_azul.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	return &Singleplayer{
		assets:       assets,
		player1:      player.New(10*6, 80*6, 10*6, 35*6, color.RGBA{255, 0, 0, 255}),
		enemy:        enemy.New(300*6, 80*6, 10*6, 35*6, color.RGBA{0, 0, 255, 255}),
		ball:         ball.New(),
		pointsPlayer: 0,
		enemyPlayer:  0,
		background:   background,
	}
}

func (s *Singleplayer) Draw(screen *ebiten.Image) {
	// Draw Background
	backOp := ebiten.GeoM{}
	backOp.Translate(-10, -10)
	backOp.Scale(2, 2)
	screen.DrawImage(s.background, &ebiten.DrawImageOptions{GeoM: backOp})

	// Draw Player 1
	s.player1.Draw(screen)

	// Draw Enemy
	s.enemy.Draw(screen)

	// Draw Ball
	s.ball.Draw(screen)
}

func (s *Singleplayer) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		s.player1.PosY -= 10
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		s.player1.PosY += 10
	}

	s.ball.Update()

	// Enemy logic
	s.enemy.BallY = s.ball.PosY
	s.enemy.Update()

	return nil
}
