package scenes

import "github.com/hajimehoshi/ebiten/v2"

type Scenes interface {
	Update() error
	Draw(*ebiten.Image)
}
