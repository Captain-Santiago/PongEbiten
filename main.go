package main

import (
	"embed"
	"errors"
	"fmt"
	"log"

	"github.com/Captain-Santiago/PongEbiten/config"
	"github.com/Captain-Santiago/PongEbiten/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var AssetServer embed.FS

var GameSingleton *Game

type Game struct {
	SceneManager            *scenes.SceneManager
	GameConfig              *config.GameConfig
	FrameCounter            uint
	SecondsSinceGameStarted uint
	musicPlayerCh           chan *AudioPlayer
	errCh                   chan error
}

// Run 60 times a second
func (g *Game) Update() error {
	// Start Audio Player
	// To do

	// Check game configs
	if ebiten.IsKeyPressed(ebiten.KeyF11) {
		g.GameConfig.ToggleFullscreen()
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		return errors.New("close_game")
	}

	// Counting seconds since game start
	g.FrameCounter++
	if g.FrameCounter >= 60 {
		g.FrameCounter = 0
		g.SecondsSinceGameStarted++
		fmt.Printf("Seconds passed: %d\n", g.SecondsSinceGameStarted)
	}
	return nil
}

// Vsynced, cannot be predicted
func (g *Game) Draw(screen *ebiten.Image) {
	// Receive screen reference
	g.SceneManager.CurrentRunningScene(screen)

	// Test if the game changed its config
	if g.GameConfig.Fullscreen != ebiten.IsFullscreen() {
		ebiten.SetFullscreen(g.GameConfig.Fullscreen)
	}

	if g.SecondsSinceGameStarted >= 5 {
		g.SceneManager.StartTitleScreen()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 4, outsideHeight / 4
}

func main() {
	gamecfg := config.GameConfig{Width: 1280, Height: 720, Title: "Go Pong Go!!!", Fullscreen: false}

	ebiten.SetWindowSize(gamecfg.Width, gamecfg.Height)
	ebiten.SetWindowTitle(gamecfg.Title)
	ebiten.SetVsyncEnabled(true)

	GameSingleton = &Game{}
	GameSingleton.SceneManager = scenes.CreateSceneManager(&AssetServer)
	GameSingleton.GameConfig = &gamecfg

	if err := ebiten.RunGame(GameSingleton); err != nil {
		if err.Error() != "close_game" {
			log.Fatal(err)
		}

	}
}
