package config

// Configuration represents the system configuration.
type Configuration struct {
	Debug        bool
	Verbose      bool
	CustomConfig interface{}
}
