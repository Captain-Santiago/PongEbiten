package singleplayer

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	"log"

	"github.com/Captain-Santiago/PongEbiten/scenes/ball"
	"github.com/Captain-Santiago/PongEbiten/scenes/player"
	"github.com/hajimehoshi/ebiten/v2"
)

type Singleplayer struct {
	assets *embed.FS

	player1      *player.Player
	enemy        *player.Player
	ball         *ball.Ball
	pointsPlayer int
	enemyPlayer  int
	background   *ebiten.Image
}

func New(assets *embed.FS) *Singleplayer {
	backarray, err := assets.ReadFile("assets/game/ceu_azul.jpeg")
	if err != nil {
		log.Fatalln(err)
	}

	backimg, _, err := image.Decode(bytes.NewBuffer(backarray))
	if err != nil {
		log.Fatalln(err)
	}

	background := ebiten.NewImageFromImage(backimg)

	return &Singleplayer{
		assets:       assets,
		player1:      player.NewPlayer(10*6, 80*6, 10*6, 35*6, color.RGBA{255, 0, 0, 255}),
		enemy:        player.NewPlayer(300*6, 80*6, 10*6, 35*6, color.RGBA{0, 0, 255, 255}),
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

	return nil
}
