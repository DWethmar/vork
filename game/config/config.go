package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Config is the configuration for the game.
type Config struct {
	DBPath     string    `json:"db_path"`
	SaveName   string    `json:"save_name"`
	CreatedAt  time.Time `json:"created_at"`
	saveFolder string    // Internal field, not exported or serialized
}

// New creates a new Config struct with default values.
func New(saveName string, parentFolder string) *Config {
	// Compute the save folder path
	sf := filepath.Join(parentFolder, saveName)
	return &Config{
		DBPath:     filepath.Join(sf, "game.db"),
		SaveName:   saveName,
		CreatedAt:  time.Now(),
		saveFolder: sf,
	}
}

// Save writes the Config struct's data to a JSON file.
func (c *Config) Save() error {
	// Ensure the save folder exists
	if err := os.MkdirAll(c.saveFolder, 0755); err != nil {
		return err
	}

	// Marshal the Config struct into JSON format
	data, err := json.MarshalIndent(c, "", "  ") // Indent for readability.
	if err != nil {
		return err
	}

	// Construct the full path to the config file
	configFilePath := filepath.Join(c.saveFolder, "config.json")

	// Write the JSON data to the file
	if err = os.WriteFile(configFilePath, data, 0600); err != nil {
		return err
	}

	return nil
}

// Load reads a JSON file and unmarshals its content into a Config struct.
func Load(saveName string, parentFolder string) (*Config, error) {
	// Construct the full path to the save folder and config file
	saveFolder := filepath.Join(parentFolder, saveName)
	configFilePath := filepath.Join(saveFolder, "config.json")

	// Read the file content
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data into Config struct
	var cfg Config
	if err = json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// Set the internal saveFolder field
	cfg.saveFolder = saveFolder

	return &cfg, nil
}

// Exists checks if a save with the given name exists.
func Exists(saveName string, parentFolder string) bool {
	saveFolder := filepath.Join(parentFolder, saveName)
	configFilePath := filepath.Join(saveFolder, "config.json")
	_, err := os.Stat(configFilePath)
	return !os.IsNotExist(err)
}

// ListSaves returns a list of save names in the given folder.
func ListSaves(saveFolder string) []Config {
	// Open the save folder
	dir, err := os.ReadDir(saveFolder)
	if err != nil {
		return nil
	}

	// Iterate over the files and directories
	var saves []Config
	for _, d := range dir {
		if d.IsDir() {
			// Check if a config file exists
			configFilePath := filepath.Join(saveFolder, d.Name(), "config.json")
			_, err = os.Stat(configFilePath)
			if !os.IsNotExist(err) {
				// Load the config file
				cfg, lErr := Load(d.Name(), saveFolder)
				if lErr == nil {
					saves = append(saves, *cfg)
				}
			}
		}
	}

	return saves
}

// Delete removes a save folder and its contents.
func Delete(saveName string, parentFolder string) error {
	saveFolder := filepath.Join(parentFolder, saveName)
	return os.RemoveAll(saveFolder)
}
