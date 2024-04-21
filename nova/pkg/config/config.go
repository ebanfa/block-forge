package config

import "time"

// BlockchainConfig represents the configuration for the entire blockchain system.
type BlockchainConfig struct {
	Version      string         `json:"version"` // Version of the configuration
	Name         string         `json:"name"`
	Frontends    []string       `json:"frontends"`
	Coins        []string       `json:"coins"`
	AbciHandlers []string       `json:"abciHandlers"`
	Modules      []ModuleConfig `json:"modules"`
	Security     SecurityConfig `json:"security"`
	Logging      LoggingConfig  `json:"logging"`
}

// ModuleConfig represents the configuration for a module within the blockchain system.
type ModuleConfig struct {
	Name             string          `json:"name"`
	Version          string          `json:"version"` // Version of the module
	Ibc              bool            `json:"ibc"`
	Dependencies     []string        `json:"dependencies"`
	EntityConfigDir  string          `json:"entityConfigDir"`
	MessageConfigDir string          `json:"messageConfigDir"`
	QueryConfigDir   string          `json:"queryConfigDir"`
	Entities         []EntityConfig  // Loaded entity configurations
	Messages         []MessageConfig // Loaded message configurations
	Queries          []QueryConfig   // Loaded query configurations
}

// EntityConfig represents the configuration for an entity within a module.
type EntityConfig struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Fields      []FieldConfig `json:"fields"`
}

// FieldConfig represents the configuration for a field within an entity.
type FieldConfig struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// MessageConfig represents the configuration for a message within a module.
type MessageConfig struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Fields      []FieldConfig `json:"fields"`
	Response    struct {
		Fields []FieldConfig `json:"fields"`
	} `json:"response"`
}

// QueryConfig represents the configuration for a query within a module.
type QueryConfig struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Fields      []FieldConfig `json:"fields"`
}

// SecurityConfig represents the security-related configuration options.
type SecurityConfig struct {
	EncryptionKey       string   `json:"encryptionKey"`
	EncryptionAlgorithm string   `json:"encryptionAlgorithm"` // Flexibility for different encryption algorithms
	EncryptionSettings  struct { // Additional encryption settings
		KeyRotationInterval time.Duration `json:"keyRotationInterval"` // Key rotation interval
		EncryptionAtRest    bool          `json:"encryptionAtRest"`    // Enable data encryption at rest
	} `json:"encryptionSettings"`
	KeyManagement struct { // Key management features
		KeyGenerationPolicy  string     `json:"keyGenerationPolicy"` // Policy for key generation
		KeyRotationPolicy    string     `json:"keyRotationPolicy"`   // Policy for key rotation
		KeyStorageMechanism  string     `json:"keyStorageMechanism"` // Mechanism for secure key storage
		KeyAccessControl     string     `json:"keyAccessControl"`    // Access control for encryption keys
		KeyLifecyclePolicies []struct { // Key lifecycle policies
			KeyExpirationPeriod time.Duration `json:"keyExpirationPeriod"` // Expiration period for keys
			KeyArchivalPolicy   string        `json:"keyArchivalPolicy"`   // Policy for key archival
		} `json:"keyLifecyclePolicies"`
	} `json:"keyManagement"`
	DataProtection struct { // Data encryption settings
		EncryptionAtTransit bool   `json:"encryptionAtTransit"` // Enable data encryption in transit
		HashAlgorithm       string `json:"hashAlgorithm"`       // Cryptographic hash algorithm for data integrity
	} `json:"dataProtection"`
	AccessControl struct { // Access control and authorization settings
		RoleBasedAccessControl bool     `json:"roleBasedAccessControl"` // Enable role-based access control
		AccessControlList      []string `json:"accessControlList"`      // List of authorized entities
	} `json:"accessControl"`
}

// LoggingConfig represents the logging configuration options.
type LoggingConfig struct {
	Level string `json:"level"`
}
