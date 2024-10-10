package audiomaster

import (
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const AUDIO_SAMPLE_RATE = 48000

// AudioPlayer represents the current audio state.
type AudioMaster struct {
	Context *audio.Context
	Player  *audio.Player
	Volume  float32
	mute    bool
}

func New(mute bool) *AudioMaster {
	return &AudioMaster{
		Context: audio.NewContext(AUDIO_SAMPLE_RATE),
		Volume:  0.1,
		mute:    mute,
	}
}

func (a *AudioMaster) PlaySong(io *io.Reader) {
	var err error

	if a.Player != nil && a.Player.IsPlaying() {
		a.Player.Close()
	}

	a.Player, err = a.Context.NewPlayerF32(*io)
	if err != nil {
		log.Fatalln("Could not play song:", err)
	}

	a.Player.Play()
}

func (a *AudioMaster) ToggleMute() {
	if a.Player == nil || !a.Player.IsPlaying() {
		return
	}

	if a.mute {
		a.Player.SetVolume(0)
		a.mute = true

	} else {
		a.Player.SetVolume(a.Player.Volume())
		a.mute = false

	}
}
