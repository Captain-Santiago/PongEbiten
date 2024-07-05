package logo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawLogoScreen(image *ebiten.Image) {
	ebitenutil.DebugPrint(image, "Logo Screen Test!!!")
}
