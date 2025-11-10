package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Initialized   bool      `json:"initialized"`
	RepoPath      string    `json:"repo_path"`
	CreatedAt     time.Time `json:"created_at"`
	NotesDir      string    `json:"notes_dir"`
	LastGenerated string    `json:"last_generated"`
	Version       string    `json:"version"`
}

var configPath = filepath.Join(".relica", "config.json")

func Exists() bool {
	_, err := os.Stat(configPath)
	return err == nil
}

func Save(cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0644)
}

func Load() (*Config, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}