package plugin

import (
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// Plugin represents a plugin in the system.
type PluginInterface interface {
	system.SystemServiceInterface

	// RegisterEventHandlers registers event handlers provided by the plugin.
	RegisterEventHandlers() error

	// RegisterComponents registers additional operations provided by the plugin.
	RegisterComponents() error
}
