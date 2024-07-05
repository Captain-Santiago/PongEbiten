package config

type GameConfig struct {
	Width, Height int
	Title         string
	Vsync         bool
	Fullscreen    bool
}

func (gc *GameConfig) ToggleFullscreen() {
	if gc.Fullscreen {
		gc.Fullscreen = false
	} else {
		gc.Fullscreen = true
	}
}

// To do
// func (gc *GameConfig) UpdateConfigs(newConfig *GameConfig, gameSingleton *ebiten.Game) {
// 	if gc.Width != newConfig.Width || gc.Height != newConfig.Height {

// 	}
// }
