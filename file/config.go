package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

const appPath = "app"

// gitConfig is the configuration for a git repository.
type GitConfig struct {
	AutoInit   bool `toml:"auto_init"`
	AutoCommit bool `toml:"auto_commit"`
}

// collectionConfig is the configuration for a collection.
type CollectionConfig struct {
	Git   GitConfig   `toml:"git"`
	Style StyleConfig `toml:"style"`
}

// styleConfig is the configuration for the style of the application.
type StyleConfig struct {
	PrimaryColor        string `toml:"primary_color"`
	PrimaryColorSubdued string `toml:"primary_color_subdued"`
	BrightGreenColor    string `toml:"bright_green"`
	GreenColor          string `toml:"green"`
	BrightRedColor      string `toml:"bright_red"`
	RedColor            string `toml:"red"`
	ForegroundColor     string `toml:"foreground"`
	BackgroundColor     string `toml:"background"`
	GrayColor           string `toml:"gray"`
	BlackColor          string `toml:"black"`
	WhiteColor          string `toml:"white"`
}

// Config is the main configuration struct.
// It contains all the configuration options at the application and collection level.
type Config struct {
	Collection CollectionConfig `toml:"collection"`
}

func GetConfig(collectionPath string) (*Config, error) {
	cfg := DefaultConfig()

	// Read app config file
	// TODO: Replace with a constant for the config file name
	file, err := os.Open(filepath.Join(appPath, ".barber.toml"))
	if err != nil {
		return nil, fmt.Errorf("Error opening config file: %s", err)
	}
	defer file.Close()

	if _, err := toml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, fmt.Errorf("Error decoding config file: %s", err)
	}

	// TODO: Read collection config file
	file, err = os.Open(collectionPath)
	if err != nil {
		return nil, fmt.Errorf("Error opening config file: %s", err)
	}
	defer file.Close()

	if _, err := toml.NewDecoder(file).Decode(cfg); err != nil {
		return nil, fmt.Errorf("Error decoding config file: %s", err)
	}

	return cfg, nil
}

// TODO: Make private after fixing config functionality
func DefaultConfig() *Config {
	// Set default values
	return &Config{
		Collection: CollectionConfig{
			Git: GitConfig{
				AutoInit:   true,
				AutoCommit: true,
			},
			Style: StyleConfig{ // TODO: Revisit the default colors
				PrimaryColor:        "#AFBEE1",
				PrimaryColorSubdued: "#64708D",
				BrightGreenColor:    "#BCE1AF",
				GreenColor:          "#527251",
				BrightRedColor:      "#E49393",
				RedColor:            "#A46060",
				ForegroundColor:     "15",
				BackgroundColor:     "235",
				GrayColor:           "241",
				BlackColor:          "#373b41",
				WhiteColor:          "#FFFFFF",
			},
		},
	}
}
