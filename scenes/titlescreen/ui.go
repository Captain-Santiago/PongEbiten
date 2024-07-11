package titlescreen

import (
	"bytes"
	"embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	// Fonts used
	titlescreenFont *text.GoTextFaceSource
	btnFont         text.Face

	// Button UI load
	btnIdle  *ebiten.Image
	btnHover *ebiten.Image
)

func loadUI(assets *embed.FS) {
	// Loading just once in memory
	if btnIdle != nil {
		return
	}

	btnIdleBytes, err := assets.ReadFile("assets/titlescreen/button_idle.png")
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err := image.Decode(bytes.NewBuffer(btnIdleBytes))
	btnIdle = ebiten.NewImageFromImage(img)

	btnHoverBytes, err := assets.ReadFile("assets/titlescreen/button_hover.png")
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err = image.Decode(bytes.NewBuffer(btnHoverBytes))
	btnHover = ebiten.NewImageFromImage(img)
}
