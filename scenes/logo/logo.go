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

var (
	ctx                 *audio.Context
	isMusicPlaying      bool
	isScreenFullyLoaded bool
	logoScreen          *ebiten.Image
	logoEbiten          *ebiten.Image
)

func init() {
	ctx = audio.NewContext(48000)
	isMusicPlaying = false
	isScreenFullyLoaded = false
}

func Draw(screen *ebiten.Image, game_assets *embed.FS) {
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
		logoScreenDecoded, _, err := image.Decode(bytes.NewBuffer(logoImageFS))
		if err != nil {
			log.Fatalln(err)
		}
		logoScreen = ebiten.NewImageFromImage(logoScreenDecoded)

		// Get ebiten logo
		logoEbitenFS, err := game_assets.ReadFile("assets/logo/logo_ebiten.png")
		if err != nil {
			log.Fatalln(err)
		}
		// Decode byte array
		logoEbitenDecoded, _, err := image.Decode(bytes.NewBuffer(logoEbitenFS))
		if err != nil {
			log.Fatalln(err)
		}
		logoEbiten = ebiten.NewImageFromImage(logoEbitenDecoded)

		// Load the image just once
		isScreenFullyLoaded = true
	}

	// Scaling sizes
	geom := ebiten.GeoM{}
	geom.Scale(0.25, 0.25)

	geomEbiten := ebiten.GeoM{}
	geomEbiten.Scale(0.25, 0.25)
	geomEbiten.Translate(260, 120)

	screen.Fill(color.White)
	screen.DrawImage(logoScreen, &ebiten.DrawImageOptions{GeoM: geom})
	screen.DrawImage(logoEbiten, &ebiten.DrawImageOptions{GeoM: geomEbiten})
}
