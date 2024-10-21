package gameplay

import (
	"os"
	"path/filepath"

	"github.com/dwethmar/vork/config"
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
func LoadConfig(saveName, savesFolder string) (*config.Config, error) {
	var cfg *config.Config
	var err error
	// Check if the save exists
	if config.Exists(saveName, savesFolder) {
		// Load the existing config
		cfg, err = config.Load(saveName, savesFolder)
		if err != nil {
			return nil, err
		}
	} else {
		// Create a new config
		cfg = config.New(saveName, savesFolder)
		if err = cfg.Save(); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}
