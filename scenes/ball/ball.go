package ball

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	PosX, PosY float32
	radius     float32
	speed      int
}

func New() *Ball {
	return &Ball{
		PosX:   1920 / 2,
		PosY:   1080 / 2,
		radius: 50,
		speed:  10,
	}
}

func (b *Ball) Update() error {
	b.PosX += 10
	b.PosY += 10

	if b.PosX > 1920 {
		b.PosX = 0
		b.PosY = 0
	}

	return nil
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, b.PosX, b.PosY, b.radius, color.RGBA{120, 180, 100, 255}, false)
}
