package scenes

import (
	"embed"

	"github.com/Captain-Santiago/PongEbiten/scenes/game"
	"github.com/Captain-Santiago/PongEbiten/scenes/logo"
	"github.com/Captain-Santiago/PongEbiten/scenes/titlescreen"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	LOGO = iota
	TITLE_SCREEN
	SINGLEPLAYER
	MULTIPLAYER
)

type SceneManager struct {
	CurrentScene   uint8
	AssetServer    *embed.FS
	MoveX1, MoveY1 int
	MoveX2, MoveY2 int
}

func CreateSceneManager(assets *embed.FS) *SceneManager {
	return &SceneManager{CurrentScene: LOGO, AssetServer: assets}
}

func (sm *SceneManager) CurrentRunningScene(screen *ebiten.Image) {
	// Debug
	// game.Start(screen, true)
	// return

	switch sm.CurrentScene {
	case LOGO:
		logo.Draw(screen, sm.AssetServer)
	case TITLE_SCREEN:
		titlescreen.Draw(screen, sm.AssetServer)
	case SINGLEPLAYER:
		game.Start(screen, true)
	case MULTIPLAYER:
		game.Start(screen, false)
	default:
		panic("not reached")
	}
}

func (sm *SceneManager) StartTitleScreen() {
	sm.CurrentScene = TITLE_SCREEN
}

func (sm *SceneManager) StartSinglePlayer() {
	sm.CurrentScene = SINGLEPLAYER
}
