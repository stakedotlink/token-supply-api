package app

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yamlv3"
)

// Config object
type Config struct {
	Port                 int      `mapstructure:"port"`
	RPCURL               string   `mapstructure:"rpcURL"`
	Ticker               string   `mapstructure:"ticker"`
	TokenAddress         string   `mapstructure:"tokenAddress"`
	TokenMintBlockNumber uint64   `mapstructure:"tokenMintBlockNumber"`
	ExcludedAddresses    []string `mapstructure:"excludedAddresses"`
	SDLPool              string   `mapstructure:"sdlPool"`
	VestingAddresses     []string `mapstructure:"vestingAddresses"`
}

// LoadConfig returns the initialised config from file
func LoadConfig(path string) (*Config, error) {
	config.WithOptions(config.ParseEnv)
	config.AddDriver(yamlv3.Driver)

	if err := config.LoadFiles(path); err != nil {
		return nil, err
	}

	cfg := Config{}
	return &cfg, config.Decode(&cfg)
}
