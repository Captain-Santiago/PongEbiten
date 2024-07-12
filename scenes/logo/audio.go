package logo

import (
	"bytes"
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

func playLogoMusic(game_assets *embed.FS) {
	isMusicPlaying = true

	// Read music file from disk
	musicByteArray, err := game_assets.ReadFile("assets/logo/background_song.wav")
	if err != nil {
		log.Fatalln(err)
	}

	// Decode music file
	musicDecoded, err := wav.DecodeWithoutResampling(bytes.NewReader(musicByteArray))
	if err != nil {
		log.Fatalln(err)
	}

	player, err := ctx.NewPlayer(musicDecoded)
	if err != nil {
		log.Fatalln(err)
	}

	// Play music in a different green thread
	go player.Play()
}
