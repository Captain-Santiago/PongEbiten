package titlescreen

import (
	"embed"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	// Fonts used
	titlescreenFont *text.GoTextFaceSource
	btnFont         *text.Face
	isFontsLoaded   bool

	// Load UI once in memory
	ui         *ebitenui.UI
	isUILoaded bool
)

var BtnFont *text.Face

func loadUI(assets embed.FS) {
	container := ui.Container

	// Creating buttons
	startGameBtn := widget.NewButton( // specify the button's text, the font face, and the color
		widget.ButtonOpts.Text("Hello, World!", *BtnFont, &widget.ButtonTextColor{
			Idle: color.RGBA{0xdf, 0xf4, 0xff, 0xff},
		}),

		// specify that the button's text needs some padding for correct display
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  30,
			Right: 30,
		}))

	container.AddChild(startGameBtn)

	isUILoaded = true
}

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}
