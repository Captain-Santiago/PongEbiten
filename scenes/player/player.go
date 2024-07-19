package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	PosX, PosY    float32
	width, height float32
	color         color.Color
}

func NewPlayer(initialPosX, initialPosY, width, height float32, playerColor color.Color) *Player {
	return &Player{
		PosX:   initialPosX,
		PosY:   initialPosY,
		width:  width,
		height: height,
		color:  playerColor,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.PosX, p.PosY, p.width, p.height, p.color, false)
}
