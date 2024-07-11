package main

import (
	"embed"
	"errors"
	"log"

	"github.com/Captain-Santiago/PongEbiten/audio"
	"github.com/Captain-Santiago/PongEbiten/config"
	"github.com/Captain-Santiago/PongEbiten/savegame"
	"github.com/Captain-Santiago/PongEbiten/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed assets/*
var AssetServer embed.FS

type Game struct {
	SceneManager            *scenes.SceneManager
	GameConfig              *config.GameConfig
	FrameCounter            uint
	SecondsSinceGameStarted uint
	MusicPlayerCh           chan *audio.AudioPlayer
	ErrCh                   chan error
}

// Run 60 times a second
func (g *Game) Update() error {
	// Update Audio Player
	// To do

	// Check game configs
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		g.GameConfig.ToggleFullscreen()
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New(CLOSE_GAME_STR)
	}

	// Test if the game changed its config
	if g.GameConfig.Fullscreen != ebiten.IsFullscreen() {
		ebiten.SetFullscreen(g.GameConfig.Fullscreen)
	}

	// Counting seconds since game start
	g.FrameCounter++
	if g.FrameCounter >= 60 {
		g.FrameCounter = 0
		g.SecondsSinceGameStarted++
	}
	return nil
}

// Vsynced, cannot be predicted
func (g *Game) Draw(screen *ebiten.Image) {
	g.SceneManager.CurrentRunningScene(screen)

	if g.SecondsSinceGameStarted >= 4 {
		g.SceneManager.StartTitleScreen()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
}

func main() {
	gamecfg := config.NewGameConfig()
	savegame.InitSaveFile(gamecfg.SaveFilePath)

	ebiten.SetWindowSize(gamecfg.Width, gamecfg.Height)
	ebiten.SetWindowTitle(gamecfg.Title)
	ebiten.SetVsyncEnabled(true)

	game := &Game{}
	game.SceneManager = scenes.CreateSceneManager(&AssetServer)
	game.GameConfig = gamecfg

	if err := ebiten.RunGame(game); err != nil {
		if err.Error() != CLOSE_GAME_STR {
			log.Fatal(err)
		}

	}

	savegame.CloseSaveFile()
}
