package titlescreen

import (
	"embed"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type TitleScreen struct {
	assets *embed.FS

	IsSingleplayer bool
	IsMultiplayer  bool

	AudioContext *audio.Context
	AudioPlayer  *audio.Player
}

func New(assets *embed.FS, audioCtx *audio.Context) *TitleScreen {
	t := TitleScreen{
		assets:         assets,
		IsSingleplayer: false,
		IsMultiplayer:  false,
		AudioContext:   audioCtx,
		AudioPlayer:    loadBackgroundSound(assets, audioCtx),
	}

	readFonts(assets)
	loadUI(assets)

	return &t
}

func (t *TitleScreen) Update() error {
	if !t.AudioPlayer.IsPlaying() {
		if err := t.AudioPlayer.Rewind(); err != nil {
			log.Fatalln("Background audio could not be rewinded:", err)
		}

		t.AudioPlayer.Play()
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		t.IsSingleplayer = true

	} else if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		t.IsMultiplayer = true

	}

	return nil
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
