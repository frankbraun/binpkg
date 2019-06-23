// Package config implements the config.binpkg configuration format.
package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"time"
)

// Config for binary packages.
type Config struct {
	URLs []string // Download URLs
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

// RandomURLs returns the URLs from cfg in random order.
func (cfg *Config) RandomURLs() ([]string, error) {
	if len(cfg.URLs) == 0 {
		return nil, errors.New("config: no URL defined")
	}
	urls := append([]string{}, cfg.URLs...)
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(urls), func(i, j int) {
		urls[i], urls[j] = urls[j], urls[i]
	})
	return urls, nil
}
