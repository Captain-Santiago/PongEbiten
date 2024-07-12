package game

import (
	"github.com/Captain-Santiago/PongEbiten/scenes/game/multiplayer"
	"github.com/Captain-Santiago/PongEbiten/scenes/game/singleplayer"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	singlePgame *singleplayer.Singleplayer
	multiPgame  *multiplayer.Multiplayer
)

type Game interface {
	Draw(*ebiten.Image)
}

func Start(screen *ebiten.Image, isSingleplayer bool) {
	if singlePgame != nil {
		singlePgame.Draw(screen)

	} else if multiPgame != nil {
		multiPgame.Draw(screen)

	}

	if isSingleplayer {
		singlePgame = singleplayer.New()
	} else {
		multiPgame = multiplayer.New()
	}
}

func Reset() {
	singlePgame = nil
	multiPgame = nil
}
