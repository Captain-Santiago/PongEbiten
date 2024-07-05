package scenes

import (
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
}

func CreateSceneManager() *SceneManager {
	return &SceneManager{CurrentScene: LOGO}
}

func (sm *SceneManager) CurrentRunningScene(screen *ebiten.Image) {
	switch sm.CurrentScene {
	case LOGO:
		logo.DrawLogoScreen(screen)
	case TITLE_SCREEN:
		titlescreen.DrawTitleScreen(screen)
	}
}

func (sm *SceneManager) StartTitleScreen() {
	sm.CurrentScene = TITLE_SCREEN
}
