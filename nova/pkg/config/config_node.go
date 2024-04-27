package config

// NodeType represents the type of a ConfigNode.
type NodeType int

const (
	ModuleNode NodeType = iota
	EntityNode
	MessageNode
	QueryNode
	FieldNode
)

// ConfigNode represents a node in the configuration tree.
type ConfigNode struct {
	Name     string
	Type     NodeType
	Value    interface{}
	Children []*ConfigNode
}

// ModuleConfig represents the configuration for a module node.
type ModuleConfig struct {
	Dependencies []string
	// Other module-specific configurations
}

// QueryConfig represents the configuration for a query or entity node.
type QueryConfig struct {
	// Other query-specific configurations
}

// TransactionConfig represents the configuration for a transaction node.
type TransactionConfig struct {
	Handler string
}

// FieldConfig represents the configuration for a field node.
type FieldConfig struct {
	Name string
	Type string
}

// TreeTraversalMode represents the mode of tree traversal.
type TreeTraversalMode int

const (
	PreOrder TreeTraversalMode = iota
	InOrder
	PostOrder
)

// ConfigNodeVisitor is an interface for visiting and processing ConfigNodes.
type ConfigNodeVisitor interface {
	VisitNode(node *ConfigNode) error
}

// ConfigTreeTraverser is an interface for traversing the configuration tree.
type ConfigTreeTraverser interface {
	Traverse(mode TreeTraversalMode, visitor ConfigNodeVisitor) error
}
