package config

// Configuration represents the system configuration.
type Configuration struct {
	ApplicationConfig ApplicationConfig
	CustomConfig      interface{}
}

// Configuration represents the system configuration.
type ApplicationConfig struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	ModuleLoadPath string   `json:"module_load_path"` // Path to load modules from
	Modules        []Module `json:"modules"`
}

// Module represents a module configuration in the configuration file.
type Module struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Path        string      `json:"path,omitempty"`
	Type        string      `json:"type,omitempty"`
	Config      interface{} `json:"config,omitempty"` // Module-specific configuration
}
