package gameplay

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dwethmar/vork/game/scenes/gameplay/config"
)

func getDefaultSaveFolder() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	saveFolder := filepath.Join(userHomeDir, ".vork", "saves")
	return saveFolder, nil
}

// Config loads or creates a config.
func LoadOrCreateConfig(saveName, savesFolder string) (*config.Config, error) {
	var cfg *config.Config
	var err error
	// Check if the save exists
	if config.Exists(saveName, savesFolder) {
		// Load the existing config
		cfg, err = config.Load(saveName, savesFolder)
		if err != nil {
			return nil, fmt.Errorf("failed to load config: %w", err)
		}
	} else {
		// Create a new config
		cfg = config.New(saveName, savesFolder)
		if err = cfg.Save(); err != nil {
			return nil, fmt.Errorf("failed to save config: %w", err)
		}
	}
	return cfg, nil
}
