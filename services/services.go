package services

import "github.com/Nyarum/lalyandia/services/config"

type Services struct {
	AuthConfig *config.AuthConfig
	GameConfig *config.GameConfig
}

func NewServices() *Services {
	return &Services{}
}

func (s *Services) LoadAuthConfig(path string) (*config.AuthConfig, error) {
	s.AuthConfig = config.NewAuthConfig()
	err := s.AuthConfig.Load(path)
	if err != nil {
		return nil, err
	}

	return s.AuthConfig, nil
}

func (s *Services) LoadGameConfig(path string) (*config.GameConfig, error) {
	s.GameConfig = config.NewGameConfig()
	err := s.GameConfig.Load(path)
	if err != nil {
		return nil, err
	}

	return s.GameConfig, nil
}
