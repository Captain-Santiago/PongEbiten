package logo

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var ctx *audio.Context
var isMusicPlaying bool
var isScreenFullyLoaded bool
var logoScreen image.Image

func init() {
	ctx = audio.NewContext(48000)
	isMusicPlaying = false
	isScreenFullyLoaded = false
}

func DrawLogoScreen(screen *ebiten.Image, game_assets *embed.FS) {
	// Start music as soon as possible
	if !isMusicPlaying {
		playLogoMusic(game_assets)
	}

	if !isScreenFullyLoaded {
		// Get logo byte array
		logoImageFS, err := game_assets.ReadFile("assets/logo/logo_screen.png")
		if err != nil {
			log.Fatalln(err)
		}

		// Decode byte array
		logoScreen, _, err = image.Decode(bytes.NewBuffer(logoImageFS))
		if err != nil {
			log.Fatalln(err)
		}

		// Load the image just once
		isScreenFullyLoaded = true
	}

	// Scaling sizes
	geom := ebiten.GeoM{}
	geom.Scale(0.25, 0.25)

	screen.Fill(color.White)
	screen.DrawImage(ebiten.NewImageFromImage(logoScreen), &ebiten.DrawImageOptions{GeoM: geom})
}
