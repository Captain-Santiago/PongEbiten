package logo

import (
	"bytes"
	"embed"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type LogoScreen struct {
	// Delta Time
	SecondsPassed uint
	ticks         uint

	// Assets
	game_assets *embed.FS
	logoScreen  *ebiten.Image
	logoEbiten  *ebiten.Image

	// Context
	ctx    *audio.Context
	Player *audio.Player
}

func New(assets *embed.FS, ctx *audio.Context) *LogoScreen {
	// Get Logo Image
	logoScreen, _, err := ebitenutil.NewImageFromFileSystem(assets, "assets/logo/logo_screen.png")
	if err != nil {
		log.Fatalln(err)
	}

	// Get ebiten logo
	logoEbiten, _, err := ebitenutil.NewImageFromFileSystem(assets, "assets/logo/logo_ebiten.png")
	if err != nil {
		log.Fatalln(err)
	}

	// Read music file from disk
	musicByteArray, err := assets.ReadFile("assets/logo/background_song.ogg")
	if err != nil {
		log.Fatalln(err)
	}

	// Decode music file
	musicDecoded, err := vorbis.DecodeWithoutResampling(bytes.NewReader(musicByteArray))
	if err != nil {
		log.Fatalln(err)
	}

	player, err := ctx.NewPlayer(musicDecoded)
	if err != nil {
		log.Fatalln(err)
	}

	err = player.Rewind()
	if err != nil {
		log.Fatalln("Could not rewind audio: ", err)
	}

	return &LogoScreen{
		SecondsPassed: 0,
		ticks:         0,
		game_assets:   assets,
		logoScreen:    logoScreen,
		logoEbiten:    logoEbiten,
		ctx:           ctx,
		Player:        player,
	}
}

func (l *LogoScreen) Update() error {
	// Start music as soon as possible
	if !l.Player.IsPlaying() {
		if err := l.Player.Rewind(); err != nil {
			log.Fatalln("Could not rewind audio: ", err)
		}
		l.Player.Play()
	}

	l.ticks += 1

	if l.ticks == 60 {
		l.ticks = 0
		l.SecondsPassed++
	}

	return nil
}

func (l *LogoScreen) Draw(screen *ebiten.Image) {
	// Scaling sizes
	geom := ebiten.GeoM{}
	geom.Scale(1.5, 1.5)

	geomEbiten := ebiten.GeoM{}
	geomEbiten.Translate(260*6, 120*6)

	screen.Fill(color.White)
	screen.DrawImage(l.logoScreen, &ebiten.DrawImageOptions{GeoM: geom})
	screen.DrawImage(l.logoEbiten, &ebiten.DrawImageOptions{GeoM: geomEbiten})
}
