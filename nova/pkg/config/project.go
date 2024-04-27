package config

// Project represents the scaffolding project.
type Project struct {
	ID              string
	Name            string
	Description     string
	ConfigTree      *ConfigTree
	DependencyGraph *DependencyGraph
	// Add other fields as needed
}

// NewProject creates a new project with the given ID, name, and description.
func NewProject(ID, name, description string) *Project {
	return &Project{
		ID:          ID,
		Name:        name,
		Description: description,
	}
}
