package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	CacheDir   string
	ShowImage  bool
	ImageWidth int
	ImageHeight int
}

func NewConfig() *Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = os.Getenv("HOME")
	}
	
	cacheDir := filepath.Join(homeDir, ".anifetch")
	
	return &Config{
		CacheDir:   cacheDir,
		ShowImage:  true,
		ImageWidth: 200,
		ImageHeight: 200,
	}
}

func (c *Config) EnsureCacheDir() error {
	return os.MkdirAll(c.CacheDir, 0755)
}

func (c *Config) GetCacheDir() string {
	return c.CacheDir
}

func (c *Config) SetShowImage(show bool) {
	c.ShowImage = show
}

func (c *Config) SetImageSize(width, height int) {
	c.ImageWidth = width
	c.ImageHeight = height
} 