package titlescreen

import (
	"embed"
	_ "image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

var (
	// Fonts used
	titlescreenFont text.Face
	btnFont         text.Face

	// Button UI load
	btnIdle  *ebiten.Image
	btnHover *ebiten.Image
)

func loadUI(assets *embed.FS) {
	var err error
	btnIdle, _, err = ebitenutil.NewImageFromFileSystem(assets, "assets/titlescreen/button_idle.png")
	if err != nil {
		log.Fatalln(err)
	}

	btnHover, _, err = ebitenutil.NewImageFromFileSystem(assets, "assets/titlescreen/button_hover.png")
	if err != nil {
		log.Fatalln(err)
	}
}

func readFonts(assets *embed.FS) {
	// Title font load
	title_font, err := assets.ReadFile("assets/fonts/kidpixies/KidpixiesRegular-p0Z1.ttf")
	if err != nil {
		log.Fatalln(err)
	}

	s, err := truetype.Parse(title_font)
	if err != nil {
		log.Fatal(err)
	}
	title_font_Face := truetype.NewFace(s, &truetype.Options{
		Size:    8 * 6,
		DPI:     300,
		Hinting: font.HintingFull,
	})
	titlescreenFont = text.NewGoXFace(title_font_Face)

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
		Size:    5,
		DPI:     300,
		Hinting: font.HintingFull,
	})
	btnFont = text.NewGoXFace(btnFontFace)
}
