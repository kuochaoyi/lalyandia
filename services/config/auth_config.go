package config

type AuthConfig struct {
	Base
}

func NewAuthConfig() *AuthConfig {
	return &AuthConfig{}
}

func (ac *AuthConfig) Load(path string) error {
	return nil
}
