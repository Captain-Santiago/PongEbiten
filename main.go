package main

import (
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/Captain-Santiago/PongEbiten/audiomaster"
	"github.com/Captain-Santiago/PongEbiten/config"
	"github.com/Captain-Santiago/PongEbiten/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const CLOSE_GAME_STR = "close_game"

//go:embed assets/*
var AssetServer embed.FS

type Game struct {
	SceneManager            *scenes.SceneManager
	GameConfig              *config.GameConfig
	FrameCounter            uint
	SecondsSinceGameStarted uint

	AudioMaster *audiomaster.AudioMaster
}

// Run 60 times a second
func (g *Game) Update() error {
	// Check game configs
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.GameConfig.ToggleFullscreen()

		// Quit game
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New(CLOSE_GAME_STR)

		// Toggle mute music if requested
	} else if inpututil.IsKeyJustPressed(ebiten.KeyM) {
		g.AudioMaster.ToggleMute()

	}

	// Test if the game changed its config
	if g.GameConfig.Fullscreen != ebiten.IsFullscreen() {
		ebiten.SetFullscreen(g.GameConfig.Fullscreen)
	}

	g.SceneManager.Update()

	return nil
}

// Vsynced, cannot be predicted
func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.CurrentScene.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.ActualFPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1920, 1080
}

func main() {
	gamecfg := config.New()

	ebiten.SetWindowSize(gamecfg.Width, gamecfg.Height)
	ebiten.SetWindowTitle(gamecfg.Title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(true)

	gameaudio := audiomaster.New(gamecfg.AudioMute)

	game := &Game{
		SceneManager: scenes.New(&AssetServer, gameaudio),
		GameConfig:   gamecfg,
		AudioMaster:  gameaudio,
	}

	if err := ebiten.RunGame(game); err != nil {
		if err.Error() != CLOSE_GAME_STR {
			log.Fatal(err)
		}
	}
}
