package config

// DependencyGraph represents the dependency graph of the project.
type DependencyGraph struct {
	Nodes []string
	Edges map[string][]string
}

// NewDependencyGraph creates a new dependency graph.
func NewDependencyGraph(nodes []string, edges map[string][]string) *DependencyGraph {
	return &DependencyGraph{
		Nodes: nodes,
		Edges: edges,
	}
}
