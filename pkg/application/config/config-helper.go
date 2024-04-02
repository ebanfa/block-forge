package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// LoadConfigurationFromFile loads the configuration from a file at the given path.
func LoadConfigurationFromFile(filePath string, target interface{}) error {
	// Read the configuration file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read configuration file: %v", err)
	}

	// Unmarshal the JSON data into the target struct
	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal configuration data: %v", err)
	}

	return nil
}
