package multiplayer

import "github.com/hajimehoshi/ebiten/v2"

type Multiplayer struct{}

func New() *Multiplayer {
	return &Multiplayer{}
}

func (m *Multiplayer) Draw(screen *ebiten.Image) {
	// Draw player 1

	// Draw player 2
}
