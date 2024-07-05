package config

type GameConfig struct {
	Width, Height int
	Title         string
	Vsync         bool
}

// To do
func (gc *GameConfig) UpdateConfigs() {

}
