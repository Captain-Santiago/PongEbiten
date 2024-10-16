package titlescreen

import (
	"embed"
	"image/color"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
)

type TitleScreen struct {
	assets *embed.FS

	IsSingleplayer bool
	IsMultiplayer  bool
}

func New(assets *embed.FS) *TitleScreen {
	readFonts(assets)
	loadUI(assets)

	return &TitleScreen{
		assets:         assets,
		IsSingleplayer: false,
		IsMultiplayer:  false,
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

func (t *TitleScreen) Draw(screen *ebiten.Image) {
	// Draw Background
	screen.Fill(color.RGBA{0x12, 0x90, 0x8e, 1})

	// Draw Game Title
	op := &text.DrawOptions{}
	op.GeoM.Translate(60*6, 10*6)
	op.ColorScale.ScaleWithColor(color.RGBA{0xf9, 0x8f, 0x45, 255})
	text.Draw(screen, "Title Screen", titlescreenFont, op)

	// Draw Button Layout
	opBtnIcon := &ebiten.DrawImageOptions{}
	opBtnIcon.GeoM.Scale(0.88*6, 0.4*6)
	opBtnIcon.GeoM.Translate(70*6, 65*6)
	screen.DrawImage(btnIdle, opBtnIcon)

	// Draw Start Button Text
	opBtns := &text.DrawOptions{}
	opBtns.GeoM.Scale(0.75*6, 0.75*6)
	opBtns.GeoM.Translate(80*6, 80*6)
	opBtns.ColorScale.Scale(255, 105, 0, 1)
	text.Draw(screen, "Start Game", btnFont, opBtns)

	// Draw Button Layout
	opBtnQuitIcon := &ebiten.DrawImageOptions{}
	opBtnQuitIcon.GeoM.Scale(0.88*6, 0.6*6)
	opBtnQuitIcon.GeoM.Translate(70*6, 110*6)
	screen.DrawImage(btnHover, opBtnQuitIcon)

	// Draw Quit Text
	opQuitBtns := &text.DrawOptions{}
	opQuitBtns.GeoM.Scale(0.75*6, 0.75*6)
	opQuitBtns.GeoM.Translate(80*6, 120*6)
	opQuitBtns.ColorScale.Scale(255, 105, 0, 1)
	text.Draw(screen, "Quit", btnFont, opQuitBtns)
}

func (t *TitleScreen) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		t.IsSingleplayer = true

	} else if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		t.IsMultiplayer = true

	}

	return nil
}
