package titlescreen

import (
	"embed"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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
