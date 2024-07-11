package config

type GameConfig struct {
	Width, Height int
	Title         string
	Vsync         bool
	Fullscreen    bool

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

func NewGameConfig() *GameConfig {
	return &GameConfig{
		Width:        1280,
		Height:       720,
		Title:        "Go Pong Go!!!",
		Fullscreen:   false,
		SaveFilePath: "./save.dat",
	}
}

// To do
// func (gc *GameConfig) UpdateConfigs(newConfig *GameConfig, gameSingleton *ebiten.Game) {
// 	if gc.Width != newConfig.Width || gc.Height != newConfig.Height {

// 	}
// }
