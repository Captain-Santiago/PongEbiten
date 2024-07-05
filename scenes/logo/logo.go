package logo

import (
	"bytes"
	"embed"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var ctx *audio.Context
var isMusicPlaying bool

func init() {
	ctx = audio.NewContext(48000)
	isMusicPlaying = false
}

func playLogoMusic(game_assets *embed.FS) {
	isMusicPlaying = true

	// Read music file from disk
	musicByteArray, err := game_assets.ReadFile("assets/logo/background_song.wav")
	if err != nil {
		log.Fatalln(err)
	}

	// Decode music file
	musicDecoded, err := wav.DecodeWithoutResampling(bytes.NewReader(musicByteArray))
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		player, err := ctx.NewPlayer(musicDecoded)
		if err != nil {
			log.Fatalln(err)
		}

		player.Play()
	}()
}

func DrawLogoScreen(screen *ebiten.Image, game_assets *embed.FS) {
	// Start music as soon as possible
	if !isMusicPlaying {
		playLogoMusic(game_assets)
	}

	// Get logo byte array
	logoImageFS, err := game_assets.ReadFile("assets/logo/logo_screen.png")
	if err != nil {
		log.Fatalln(err)
	}

	// Decode byte array
	img, _, err := image.Decode(bytes.NewBuffer(logoImageFS))
	if err != nil {
		log.Fatalln(err)
	}

	// Scaling sizes
	geom := ebiten.GeoM{}
	geom.Scale(0.25, 0.25)

	screen.Fill(color.White)
	screen.DrawImage(ebiten.NewImageFromImage(img), &ebiten.DrawImageOptions{GeoM: geom})
}
