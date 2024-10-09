package titlescreen

import (
	"bytes"
	"embed"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

func loadBackgroundSound(assets *embed.FS, ctx *audio.Context) *audio.Player {
	var err error

	songbytes, err := assets.ReadFile("assets/titlescreen/audio/royalty.ogg")
	if err != nil {
		log.Fatalln("Could not load song:", err)
	}

	decodedsound, err := vorbis.DecodeF32(bytes.NewReader(songbytes))
	if err != nil {
		log.Fatalln("Could not decode sound:", err)
	}

	player, err := ctx.NewPlayerF32(decodedsound)
	if err != nil {
		log.Fatalln("Could not create a new player:", err)
	}

	player.SetVolume(0.1)

	printLicense(assets)
	return player
}

func printLicense(assets *embed.FS) {
	license, err := assets.ReadFile("assets/titlescreen/audio/LICENSE")
	if err != nil {
		log.Fatalln("Could not read license file:", err)
	}

	fmt.Print(string(license))
}
