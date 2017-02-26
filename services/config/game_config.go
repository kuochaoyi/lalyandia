package config

type GameConfig struct {
	Base
}

func NewGameConfig() *GameConfig {
	return &GameConfig{}
}

func (gc *GameConfig) Load(path string) error {
	return nil
}
