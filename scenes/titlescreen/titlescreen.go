package titlescreen

import (
	"bytes"
	"embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	titlescreenFont *text.GoTextFaceSource
	isFontLoaded    bool
)

func init() {
	isFontLoaded = false
}

func readFont(assets embed.FS) {
	font, err := assets.ReadFile("assets/fonts/kidpixies/KidpixiesRegular-p0Z1.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	s, err := text.NewGoTextFaceSource(bytes.NewReader(font))
	if err != nil {
		log.Fatal(err)
	}
	titlescreenFont = s

	isFontLoaded = true
}

func DrawTitleScreen(screen *ebiten.Image, assets embed.FS) {
	if !isFontLoaded {
		readFont(assets)
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(90, 10)
	op.ColorScale.ScaleWithColor(color.White)

	tface := &text.GoTextFace{Size: 32, Source: titlescreenFont}

	// Draw image on screen
	screen.Fill(color.RGBA{160, 32, 240, 0})
	text.Draw(screen, "Title Screen", tface, op)
}
