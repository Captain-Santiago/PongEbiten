package config

type GameConfig struct {
	Width, Height int
	Title         string
	Vsync         bool
	Fullscreen    bool
	AudioMute     bool

	// Save File
	SaveFilePath string
}

func (gc *GameConfig) ToggleFullscreen() {
	if gc.Fullscreen {
		gc.Fullscreen = false
	} else {
		gc.Fullscreen = true
	}
}

func New() *GameConfig {
	return &GameConfig{
		Width:      1280,
		Height:     720,
		Title:      "Go Pong Go!!!",
		Fullscreen: false,
		AudioMute:  false,
	}
}
