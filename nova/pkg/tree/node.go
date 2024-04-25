package tree

// NodeType represents the type of a ConfigNode.
type NodeType int

const (
	ModuleNode NodeType = iota
	TransactionNode
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
