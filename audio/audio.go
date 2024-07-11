package audio

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const AUDIO_SAMPLE_RATE = 48000

var AudioMasterControl AudioMaster

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

func NewPlayer(musicPlayerCh chan *AudioPlayer, errCh chan error) *AudioPlayer {
	audioContext := audio.NewContext(AUDIO_SAMPLE_RATE)

	AudioMasterControl = AudioMaster{}
	AudioMasterControl.MusicPlayerCh = make(chan *AudioPlayer)
	AudioMasterControl.ErrCh = make(chan error)

	audioP := &audio.Player{}

	return &AudioPlayer{
		AudioContext: audioContext,
		AudioPlayer:  audioP,
		Current:      0,
		Total:        0,
		Volume128:    75,
	}
}
