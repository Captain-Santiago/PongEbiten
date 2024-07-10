package scenes

import (
	"embed"

	"github.com/Captain-Santiago/PongEbiten/scenes/logo"
	"github.com/Captain-Santiago/PongEbiten/scenes/titlescreen"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	LOGO = iota
	TITLE_SCREEN
)

type SceneManager struct {
	CurrentScene uint8
	AssetServer  *embed.FS
}

func CreateSceneManager(assets *embed.FS) *SceneManager {
	return &SceneManager{CurrentScene: LOGO, AssetServer: assets}
}

func (sm *SceneManager) CurrentRunningScene(screen *ebiten.Image) {
	// Debug
	// titlescreen.DrawTitleScreen(screen, *sm.AssetServer)
	// return

	switch sm.CurrentScene {
	case LOGO:
		logo.DrawLogoScreen(screen, sm.AssetServer)
	case TITLE_SCREEN:
		titlescreen.DrawTitleScreen(screen, *sm.AssetServer)
	default:
		panic("not reached")
	}
}

func (sm *SceneManager) StartTitleScreen() {
	sm.CurrentScene = TITLE_SCREEN
}
