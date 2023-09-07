package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const appPath = "app"

// gitConfig is the configuration for a git repository.
type gitConfig struct {
	AutoInit   bool `toml:"auto_init"`
	AutoCommit bool `toml:"auto_commit"`
}

// collectionConfig is the configuration for a collection.
type collectionConfig struct {
	Git gitConfig `toml:"git"`
}

// Config is the main configuration struct.
// It contains all the configuration options at the application and collection level.
type Config struct {
	Collection collectionConfig `toml:"collection"`
}

func GetConfig(collectionPath string) *Config {
	// Set default values
	cfg := &Config{
		Collection: collectionConfig{
			Git: gitConfig{
				AutoInit:   true,
				AutoCommit: true,
			},
		},
	}

	// Read app config file
	// TODO: Replace with a constant for the config file name
	file, err := os.Open(filepath.Join(appPath, ".barber.toml"))
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}
	defer file.Close()

	if _, err := toml.NewDecoder(file).Decode(cfg); err != nil {
		log.Fatalf("Error decoding config file: %s", err)
	}

	// TODO: Read collection config file
	file, err = os.Open(collectionPath)
	if err != nil {
		log.Fatalf("Error opening config file: %s", err)
	}
	defer file.Close()

	if _, err := toml.NewDecoder(file).Decode(cfg); err != nil {
		log.Fatalf("Error decoding config file: %s", err)
	}

	return cfg
}
