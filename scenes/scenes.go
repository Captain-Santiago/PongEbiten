package scenes

import (
	"embed"

	"github.com/Captain-Santiago/PongEbiten/scenes/logo"
	"github.com/Captain-Santiago/PongEbiten/scenes/multiplayer"
	"github.com/Captain-Santiago/PongEbiten/scenes/singleplayer"
	"github.com/Captain-Santiago/PongEbiten/scenes/titlescreen"
	"github.com/hajimehoshi/ebiten/v2/audio"
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

	AudioContext *audio.Context
}

func New(assets *embed.FS) *SceneManager {
	audioCtx := audio.NewContext(48000)
	return &SceneManager{CurrentScene: logo.New(assets, audioCtx), AssetServer: assets, AudioContext: audioCtx}
}

func (sm *SceneManager) Update() error {
	switch sm.CurrentScene.(type) {
	case *logo.LogoScreen:
		if sm.CurrentScene.(*logo.LogoScreen).SecondsPassed == 4 {
			sm.CurrentScene.(*logo.LogoScreen).Player.Close()
			sm.CurrentScene = titlescreen.New(sm.AssetServer, sm.AudioContext)
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
