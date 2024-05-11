package config

import (
	"os"
	"path/filepath"
)

const (
	DataDirName      = ".nova"
	MetadataDbName   = "MetadataStore"
	MultiStoreDbName = "MultiStore"
)

type NovaConfig struct {
	UserHomeDir      string
	DatabasesDir     string
	MetadataDbName   string
	MultiStoreDbName string
}

func GetDefaultConfig() (NovaConfig, error) {
	// Get the home directory of the current user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return NovaConfig{}, err
	}

	// Define the database path within the .nova directory
	databasesDir := filepath.Join(homeDir, DataDirName, "databases")

	configuration := NovaConfig{
		UserHomeDir:      homeDir,
		DatabasesDir:     databasesDir,
		MetadataDbName:   MetadataDbName,
		MultiStoreDbName: MultiStoreDbName,
	}

	return configuration, nil
}
