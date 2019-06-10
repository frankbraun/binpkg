// Package config implements the config.binpkg configuration format.
package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	URLs []string
}

// Load a config.binpkg file from filename and return the Config struct.
func Load(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, err
}

// Marshal cfg as string.
func (cfg *Config) Marshal() string {
	jsn, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err) // should never happen
	}
	return string(jsn)
}
