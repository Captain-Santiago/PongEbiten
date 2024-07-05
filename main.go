package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/Captain-Santiago/PongEbiten/config"
	"github.com/Captain-Santiago/PongEbiten/scenes"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var AssetServer embed.FS

type Game struct {
	SceneManager            *scenes.SceneManager
	FrameCounter            uint
	SecondsSinceGameStarted uint
	musicPlayerCh           chan *AudioPlayer
	errCh                   chan error
}

// Run 60 times a second
func (g *Game) Update() error {
	// Start Audio Player
	// To do

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

	if g.SecondsSinceGameStarted >= 5 {
		g.SceneManager.StartTitleScreen()
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth / 4, outsideHeight / 4
}

func main() {
	gamecfg := config.GameConfig{Width: 1280, Height: 720, Title: "Go Pong Go!!!"}

	ebiten.SetWindowSize(gamecfg.Width, gamecfg.Height)
	ebiten.SetWindowTitle(gamecfg.Title)
	ebiten.SetVsyncEnabled(true)

	gameSingleton := &Game{}
	gameSingleton.SceneManager = scenes.CreateSceneManager(&AssetServer)

	if err := ebiten.RunGame(gameSingleton); err != nil {
		log.Fatal(err)
	}
}
