package audio

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

type AudioMaster struct {
	MusicPlayerCh chan *AudioPlayer
	ErrCh         chan error
}

// AudioPlayer represents the current audio state.
type AudioPlayer struct {
	AudioContext *audio.Context
	AudioPlayer  *audio.Player
	Current      time.Duration
	Total        time.Duration
	SeBytes      []byte
	SeCh         chan []byte
	Volume128    int
}

func New() *AudioPlayer {
	const AUDIO_SAMPLE_RATE = 48000

	return &AudioPlayer{
		AudioContext: audio.NewContext(AUDIO_SAMPLE_RATE),
		Current:      0,
		Total:        0,
	}
}
