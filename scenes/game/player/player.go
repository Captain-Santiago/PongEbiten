package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	posX, posY    float32
	width, height float32
	color         color.Color
}

func NewPlayer(initialPosX, initialPosY, width, height float32, playerColor color.Color) *Player {
	return &Player{
		posX:   initialPosX,
		posY:   initialPosY,
		width:  width,
		height: height,
		color:  playerColor,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, p.posX, p.posY, p.width, p.height, p.color, false)
}
