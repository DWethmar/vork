package gameplay

import (
	"os"
	"path/filepath"
)

func getDefaultSaveFolder() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	saveFolder := filepath.Join(userHomeDir, ".vork", "saves")
	return saveFolder, nil
}
