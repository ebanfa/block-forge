package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ReadBlockchainConfig reads and parses the blockchain configuration from the given file.
func ReadBlockchainConfig(filename string) (BlockchainConfig, error) {
	var blockchainConfig BlockchainConfig
	// Read the blockchain configuration file
	file, err := os.Open(filename)
	if err != nil {
		return BlockchainConfig{}, fmt.Errorf("failed to open blockchain config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&blockchainConfig); err != nil {
		return BlockchainConfig{}, fmt.Errorf("failed to decode blockchain config: %w", err)
	}

	// Load module configurations
	for i := range blockchainConfig.Modules {
		module := &blockchainConfig.Modules[i]
		moduleConfigFile := filepath.Join(filepath.Dir(filename), module.Name, "module.json")
		moduleConfig, err := ReadModuleConfig(moduleConfigFile)
		if err != nil {
			return BlockchainConfig{}, fmt.Errorf("failed to load module config for %s: %w", module.Name, err)
		}
		module.Version = moduleConfig.Version
		module.Dependencies = moduleConfig.Dependencies
		module.EntityConfigDir = moduleConfig.EntityConfigDir
		module.MessageConfigDir = moduleConfig.MessageConfigDir
		module.QueryConfigDir = moduleConfig.QueryConfigDir
	}

	return blockchainConfig, nil
}

// ReadModuleConfig reads and parses the module configuration from the given file,
// and loads entity, message, and query configurations based on the module's directory structure.
func ReadModuleConfig(filename string) (ModuleConfig, error) {
	var moduleConfig ModuleConfig
	// Read the module configuration file
	data, err := os.ReadFile(filename)
	if err != nil {
		return ModuleConfig{}, fmt.Errorf("failed to read module config file: %w", err)
	}
	if err := json.Unmarshal(data, &moduleConfig); err != nil {
		return ModuleConfig{}, fmt.Errorf("failed to decode module config: %w", err)
	}

	// Load entity configurations
	entityConfigDir := filepath.Join(filepath.Dir(filename), moduleConfig.EntityConfigDir)

	fmt.Printf("About to load entity config: %s\n", entityConfigDir)

	entities, err := loadEntityConfigs(entityConfigDir)
	if err != nil {
		return ModuleConfig{}, fmt.Errorf("failed to load entity configs: %w", err)
	}
	moduleConfig.Entities = entities

	// Load message configurations
	messageConfigDir := filepath.Join(filepath.Dir(filename), moduleConfig.MessageConfigDir)
	messages, err := loadMessageConfigs(messageConfigDir)
	if err != nil {
		return ModuleConfig{}, fmt.Errorf("failed to load message configs: %w", err)
	}
	moduleConfig.Messages = messages

	// Load query configurations
	queryConfigDir := filepath.Join(filepath.Dir(filename), moduleConfig.QueryConfigDir)
	queries, err := loadQueryConfigs(queryConfigDir)
	if err != nil {
		return ModuleConfig{}, fmt.Errorf("failed to load query configs: %w", err)
	}
	moduleConfig.Queries = queries

	return moduleConfig, nil
}

// loadEntityConfigs loads entity configurations from the given directory.
func loadEntityConfigs(dir string) ([]EntityConfig, error) {
	// Read the directory entries
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var entities []EntityConfig
	// Iterate over directory entries
	for _, entry := range dirEntries {
		// Skip directories
		if entry.IsDir() {
			continue
		}
		// Construct the entity file path
		entityFile := filepath.Join(dir, entry.Name())
		// Read and parse the entity configuration from the file
		entityConfig, err := ReadEntityConfig(entityFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read entity config from file %s: %w", entityFile, err)
		}
		entities = append(entities, entityConfig)
	}
	return entities, nil
}

// loadMessageConfigs loads message configurations from the given directory.
func loadMessageConfigs(dir string) ([]MessageConfig, error) {
	// Read the directory entries
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var messages []MessageConfig
	// Iterate over directory entries
	for _, entry := range dirEntries {
		// Skip directories
		if entry.IsDir() {
			continue
		}
		// Construct the message file path
		messageFile := filepath.Join(dir, entry.Name())
		// Read and parse the message configuration from the file
		messageConfig, err := ReadMessageConfig(messageFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read message config from file %s: %w", messageFile, err)
		}
		messages = append(messages, messageConfig)
	}
	return messages, nil
}

// loadQueryConfigs loads query configurations from the given directory.
func loadQueryConfigs(dir string) ([]QueryConfig, error) {
	// Read the directory entries
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var queries []QueryConfig
	// Iterate over directory entries
	for _, entry := range dirEntries {
		// Skip directories
		if entry.IsDir() {
			continue
		}
		// Construct the query file path
		queryFile := filepath.Join(dir, entry.Name())
		// Read and parse the query configuration from the file
		queryConfig, err := ReadQueryConfig(queryFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read query config from file %s: %w", queryFile, err)
		}
		queries = append(queries, queryConfig)
	}
	return queries, nil
}

// ReadEntityConfig reads and parses the entity configuration from the given file.
func ReadEntityConfig(filename string) (EntityConfig, error) {
	var entityConfig EntityConfig
	// Read the entity configuration file
	data, err := os.ReadFile(filename)
	if err != nil {
		return EntityConfig{}, fmt.Errorf("failed to read entity config file: %w", err)
	}
	// Parse the entity configuration from the file data
	if err := json.Unmarshal(data, &entityConfig); err != nil {
		return EntityConfig{}, fmt.Errorf("failed to decode entity config: %w", err)
	}
	return entityConfig, nil
}

// ReadMessageConfig reads and parses the message configuration from the given file.
func ReadMessageConfig(filename string) (MessageConfig, error) {
	var messageConfig MessageConfig
	// Read the message configuration file
	data, err := os.ReadFile(filename)
	if err != nil {
		return MessageConfig{}, fmt.Errorf("failed to read message config file: %w", err)
	}
	// Parse the message configuration from the file data
	if err := json.Unmarshal(data, &messageConfig); err != nil {
		return MessageConfig{}, fmt.Errorf("failed to decode message config: %w", err)
	}
	return messageConfig, nil
}

// ReadQueryConfig reads and parses the query configuration from the given file.
func ReadQueryConfig(filename string) (QueryConfig, error) {
	var queryConfig QueryConfig
	// Read the query configuration file
	data, err := os.ReadFile(filename)
	if err != nil {
		return QueryConfig{}, fmt.Errorf("failed to read query config file: %w", err)
	}
	// Parse the query configuration from the file data
	if err := json.Unmarshal(data, &queryConfig); err != nil {
		return QueryConfig{}, fmt.Errorf("failed to decode query config: %w", err)
	}
	return queryConfig, nil
}

/* // ReadEntityConfig reads and parses the entity configuration from the given file.

// ReadSecurityConfig reads and parses the security configuration from the given file.
func ReadSecurityConfig(filename string) (SecurityConfig, error) {
	// Similar implementation as ReadModuleConfig and others
}

// ReadLoggingConfig reads and parses the logging configuration from the given file.
func ReadLoggingConfig(filename string) (LoggingConfig, error) {
	// Similar implementation as ReadModuleConfig and others
}

// ReadModuleConfigs reads and parses module configurations from the given directory.
func ReadModuleConfigs(dir string) ([]ModuleConfig, error) {
	// Iterate over files in the directory and call ReadModuleConfig for each file
}

// ReadEntityConfigs reads and parses entity configurations from the given directory.
func ReadEntityConfigs(dir string) ([]EntityConfig, error) {
	// Similar implementation as ReadModuleConfigs
}

// ReadMessageConfigs reads and parses message configurations from the given directory.
func ReadMessageConfigs(dir string) ([]MessageConfig, error) {
	// Similar implementation as ReadModuleConfigs
}

// ReadQueryConfigs reads and parses query configurations from the given directory.
func ReadQueryConfigs(dir string) ([]QueryConfig, error) {
	// Similar implementation as ReadModuleConfigs
}
*/
