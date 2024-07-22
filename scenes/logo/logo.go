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
	ctx            *audio.Context
	isMusicPlaying bool
	logoScreen     *ebiten.Image
	logoEbiten     *ebiten.Image
)

type LogoScreen struct {
	game_assets   *embed.FS
	SecondsPassed uint
	ticks         uint
}

func New(assets *embed.FS) *LogoScreen {
	// Get logo byte array
	logoImageFS, err := assets.ReadFile("assets/logo/logo_screen.png")
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
	logoEbitenFS, err := assets.ReadFile("assets/logo/logo_ebiten.png")
	if err != nil {
		log.Fatalln(err)
	}
	// Decode byte array
	logoEbitenDecoded, _, err := image.Decode(bytes.NewBuffer(logoEbitenFS))
	if err != nil {
		log.Fatalln(err)
	}
	logoEbiten = ebiten.NewImageFromImage(logoEbitenDecoded)

	return &LogoScreen{
		game_assets:   assets,
		SecondsPassed: 0,
		ticks:         0,
	}
}

func (l *LogoScreen) Update() error {
	if ctx == nil {
		ctx = audio.NewContext(48000)

		// ---------------------- DEBUG MUSIC OFF
		isMusicPlaying = true
		// ---------------------------
	}

	l.ticks += 1

	if l.ticks == 60 {
		l.ticks = 0
		l.SecondsPassed++
	}

	return nil
}

func (l *LogoScreen) Draw(screen *ebiten.Image) {
	// Start music as soon as possible
	if !isMusicPlaying {
		playLogoMusic(l.game_assets)
	}

	// Scaling sizes
	geom := ebiten.GeoM{}
	geom.Scale(1.5, 1.5)

	geomEbiten := ebiten.GeoM{}
	geomEbiten.Translate(260*6, 120*6)

	screen.Fill(color.White)
	screen.DrawImage(logoScreen, &ebiten.DrawImageOptions{GeoM: geom})
	screen.DrawImage(logoEbiten, &ebiten.DrawImageOptions{GeoM: geomEbiten})
}
