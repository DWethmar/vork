package config_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/dwethmar/vork/game/scenes/gameplay/config"
	"github.com/google/go-cmp/cmp"
)

// TestNew tests the New function of the config package.
func TestNew(t *testing.T) {
	saveName := "test_save"
	parentFolder := "./test_saves"

	// Clean up after test
	defer os.RemoveAll(parentFolder)

	// Create a new config
	cfg := config.New(saveName, parentFolder)

	// Verify that the Config struct is initialized correctly
	expectedSaveFolder := filepath.Join(parentFolder, saveName)
	if cfg.SaveName != saveName {
		t.Errorf("Expected SaveName '%s', got '%s'", saveName, cfg.SaveName)
	}
	if cfg.DBPath != filepath.Join(expectedSaveFolder, "game.db") {
		t.Errorf("Expected DBPath '%s', got '%s'", filepath.Join(expectedSaveFolder, "game.db"), cfg.DBPath)
	}
	if cfg.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set, got zero value")
	}
}

// TestExists tests the Exists function of the config package.
func TestExists(t *testing.T) {
	saveName := "test_save"
	parentFolder := "./test_saves"

	// Clean up after test
	defer os.RemoveAll(parentFolder)

	// Initially, the save should not exist
	if config.Exists(saveName, parentFolder) {
		t.Error("Expected save not to exist")
	}

	// Create a new config and save it
	cfg := config.New(saveName, parentFolder)
	if err := cfg.Save(); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Now, the save should exist
	if !config.Exists(saveName, parentFolder) {
		t.Error("Expected save to exist")
	}
}

// TestLoad tests the Load function of the config package.
func TestLoad(t *testing.T) {
	saveName := "test_save"
	parentFolder := "./test_saves"

	// Clean up after test
	defer os.RemoveAll(parentFolder)

	// Create a new config and save it
	createdAt := time.Now().Round(time.Second)
	cfg := config.New(saveName, parentFolder)
	cfg.CreatedAt = createdAt
	if err := cfg.Save(); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Load the config
	loadedCfg, err := config.Load(saveName, parentFolder)
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify that the loaded config matches the saved config
	if cfg.SaveName != loadedCfg.SaveName {
		t.Errorf("Expected SaveName '%s', got '%s'", cfg.SaveName, loadedCfg.SaveName)
	}
	if cfg.DBPath != loadedCfg.DBPath {
		t.Errorf("Expected DBPath '%s', got '%s'", cfg.DBPath, loadedCfg.DBPath)
	}
	if !cfg.CreatedAt.Equal(loadedCfg.CreatedAt) {
		t.Errorf("Expected CreatedAt '%v', got '%v'", cfg.CreatedAt, loadedCfg.CreatedAt)
	}
}

// TestConfig_Save tests the Save method of the Config struct.
func TestConfig_Save(t *testing.T) {
	saveName := "test_save"
	parentFolder := "./test_saves"

	// Clean up after test
	defer os.RemoveAll(parentFolder)

	// Create a new config
	cfg := config.New(saveName, parentFolder)

	// Save the config
	if err := cfg.Save(); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Verify that the config file exists
	configFilePath := filepath.Join(parentFolder, saveName, "config.json")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		t.Errorf("Config file '%s' does not exist after saving", configFilePath)
	}
}

// TestListSaves tests the ListSaves function from the game package.
func TestListSaves(t *testing.T) {
	saveFolder := "./test_saves"

	// Clean up after test
	defer os.RemoveAll(saveFolder)

	// Create some save files
	saveNames := []string{"save1", "save2", "save3"}

	for _, saveName := range saveNames {
		cfg := config.New(saveName, saveFolder)
		if err := cfg.Save(); err != nil {
			t.Fatalf("Failed to save config: %v", err)
		}
	}

	// List the saves
	saves := config.ListSaves(saveFolder)

	// Verify that the list contains the expected saves
	if len(saves) != len(saveNames) {
		t.Errorf("Expected %d saves, got %d", len(saveNames), len(saves))
	}

	for i, saveName := range saveNames {
		if diff := cmp.Diff(saveName, saves[i].SaveName); diff != "" {
			t.Errorf("Save name mismatch (-want +got):\n%s", diff)
		}
	}
}

// TestDelete tests the DeleteSave function from the game package.
func TestDelete(t *testing.T) {
	saveFolder := "./test_saves"
	saveName := "test_save"

	// Clean up after test
	defer os.RemoveAll(saveFolder)

	// Create a new config and save it
	cfg := config.New(saveName, saveFolder)
	if err := cfg.Save(); err != nil {
		t.Fatalf("Failed to save config: %v", err)
	}

	// Verify that the save exists
	if !config.Exists(saveName, saveFolder) {
		t.Error("Expected save to exist")
	}

	// Delete the save
	if err := config.Delete(saveName, saveFolder); err != nil {
		t.Fatalf("Failed to delete save: %v", err)
	}

	// Verify that the save no longer exists
	if config.Exists(saveName, saveFolder) {
		t.Error("Expected save not to exist after deletion")
	}
}
