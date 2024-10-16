package scenes

import (
	"embed"

	"github.com/Captain-Santiago/PongEbiten/scenes/logo"
	"github.com/Captain-Santiago/PongEbiten/scenes/multiplayer"
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

	case *titlescreen.TitleScreen:
		if sm.CurrentScene.(*titlescreen.TitleScreen).IsSingleplayer {
			sm.CurrentScene = singleplayer.New(sm.AssetServer)

		} else if sm.CurrentScene.(*titlescreen.TitleScreen).IsMultiplayer {
			sm.CurrentScene = multiplayer.New(sm.AssetServer)

		}
	}

	err := sm.CurrentScene.Update()
	return err
}
