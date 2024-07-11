package titlescreen

import (
	"bytes"
	"embed"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

func readFonts(assets *embed.FS) {
	// Title font load
	title_font, err := assets.ReadFile("assets/fonts/kidpixies/KidpixiesRegular-p0Z1.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	s, err := text.NewGoTextFaceSource(bytes.NewReader(title_font))
	if err != nil {
		log.Fatal(err)
	}
	titlescreenFont = s

	// Load Button Fonts
	btn_font, err := assets.ReadFile("assets/fonts/PressStart2p/PressStart2P-vaV7.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	ttfFont, err := truetype.Parse(btn_font)
	if err != nil {
		log.Fatalln(err)
	}

	btnFontFace := truetype.NewFace(ttfFont, &truetype.Options{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	btnFont = text.NewGoXFace(btnFontFace)
}

func DrawTitleScreen(screen *ebiten.Image, assets *embed.FS) {
	if titlescreenFont == nil {
		readFonts(assets)
	}

	op := &text.DrawOptions{}
	op.GeoM.Translate(70, 10)
	op.ColorScale.ScaleWithColor(color.White)

	tface := &text.GoTextFace{Size: 32, Source: titlescreenFont}

	screen.Fill(color.RGBA{160, 32, 240, 0})
	text.Draw(screen, "Title Screen", tface, op)

	// Draw buttons on screen
	loadUI(assets)
}
