package titlescreen

import (
	"bytes"
	"embed"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

func init() {
	isFontsLoaded = false
	isUILoaded = false
}

func readFonts(assets embed.FS) {
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
	BtnFont = btnFontFace

	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	isFontsLoaded = true
}

func DrawTitleScreen(screen *ebiten.Image, assets embed.FS) {
	if !isFontsLoaded {
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
