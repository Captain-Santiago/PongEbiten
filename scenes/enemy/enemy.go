package enemy

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Enemy struct {
	PosX, PosY    float32
	width, height float32
	color         color.Color

	// Follow ball logic
	BallY float32
}

func New(initialPosX, initialPosY, width, height float32, enemyColor color.Color) *Enemy {
	return &Enemy{
		PosX:   initialPosX,
		PosY:   initialPosY,
		width:  width,
		height: height,
		color:  enemyColor,
	}
}

func (p *Enemy) Update() error {
	if p.PosY < p.BallY {
		p.PosY += 10

	} else if p.PosY > p.BallY {
		p.PosY -= 10

	}

	if p.PosY > 870 {
		p.PosY = 870
	} else if p.PosY < 0 {
		p.PosY = 0
	}

	return nil
}

func (p *Enemy) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.PosX, p.PosY, p.width, p.height, p.color, false)
}
