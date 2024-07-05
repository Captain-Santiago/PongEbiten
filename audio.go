package main

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

const AUDIO_SAMPLE_RATE = 48000

// AudioPlayer represents the current audio state.
type AudioPlayer struct {
	game         *Game
	audioContext *audio.Context
	audioPlayer  *audio.Player
	current      time.Duration
	total        time.Duration
	seBytes      []byte
	seCh         chan []byte
	volume128    int
}

func NewGame() (*Game, error) {
	audioContext := audio.NewContext(AUDIO_SAMPLE_RATE)

	g := &Game{
		musicPlayerCh: make(chan *AudioPlayer),
		errCh:         make(chan error),
	}

	_, err := NewPlayer(g, audioContext)
	if err != nil {
		return nil, err
	}

	// g.musicPlayer = m
	return g, nil
}

func NewPlayer(game *Game, audioContext *audio.Context) (*audio.Player, error) {
	audioP := &audio.Player{}

	return audioP, nil
}
