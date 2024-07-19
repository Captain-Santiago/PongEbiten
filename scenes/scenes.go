package scenes

import (
	"embed"

	"github.com/Captain-Santiago/PongEbiten/scenes/logo"
	"github.com/Captain-Santiago/PongEbiten/scenes/singleplayer"
	"github.com/Captain-Santiago/PongEbiten/scenes/titlescreen"
)

const (
	LOGO = iota
	TITLE_SCREEN
	SINGLEPLAYER
	MULTIPLAYER
)

type SceneManager struct {
	CurrentScene Scenes
	AssetServer  *embed.FS
}

func New(assets *embed.FS) *SceneManager {
	return &SceneManager{CurrentScene: logo.New(assets), AssetServer: assets}
}

func (sm *SceneManager) Update() error {
	switch sm.CurrentScene.(type) {
	case *logo.LogoScreen:
		if sm.CurrentScene.(*logo.LogoScreen).SecondsPassed == 4 {
			sm.CurrentScene = titlescreen.New(sm.AssetServer)
		}

		sm.CurrentScene.Update()
	case *titlescreen.TitleScreen:
		if sm.CurrentScene.(*titlescreen.TitleScreen).NextScreen {
			sm.CurrentScene = singleplayer.New(sm.AssetServer)
		}

		sm.CurrentScene.Update()
	case *singleplayer.Singleplayer:
		sm.CurrentScene.Update()
	default:
		panic("not reached")
	}

	return nil
}
