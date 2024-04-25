package tree

import "fmt"

// ConfigTree represents the root of the configuration tree.
type ConfigTree struct {
	Root *ConfigNode
}

// NewConfigTree creates a new ConfigTree instance with an empty root node.
func NewConfigTree() *ConfigTree {
	root := &ConfigNode{}
	return &ConfigTree{
		Root: root,
	}
}

// AddNode adds a new node to the configuration tree.
func (tree *ConfigTree) AddNode(node *ConfigNode) {
	// If the tree is empty, set the node as the root.
	if tree.Root == nil {
		tree.Root = node
	} else {
		// Otherwise, add the node as a child of the root.
		tree.Root.Children = append(tree.Root.Children, node)
	}
}

// FindNode finds a node in the configuration tree based on its name.
func (tree *ConfigTree) FindNode(name string) *ConfigNode {
	return tree.findNodeRecursive(tree.Root, name)
}

func (tree *ConfigTree) findNodeRecursive(node *ConfigNode, name string) *ConfigNode {
	// If the current node is nil, return nil.
	if node == nil {
		return nil
	}
	// If the current node's name matches the target name, return the node.
	if node.Name == name {
		return node
	}
	// Recursively search through the children of the current node.
	for _, child := range node.Children {
		if found := tree.findNodeRecursive(child, name); found != nil {
			return found
		}
	}
	// If not found, return nil.
	return nil
}

// RemoveNode removes a node from the configuration tree.
func (tree *ConfigTree) RemoveNode(name string) {
	if tree.Root == nil {
		return
	}
	// Recursively remove the node from the tree.
	tree.Root = tree.removeNodeRecursive(tree.Root, name)
}

func (tree *ConfigTree) removeNodeRecursive(node *ConfigNode, name string) *ConfigNode {
	if node == nil {
		return nil
	}
	filteredChildren := []*ConfigNode{}
	// Filter out the node with the target name from the children.
	for _, child := range node.Children {
		if child.Name != name {
			filteredChildren = append(filteredChildren, child)
		}
	}
	// Set the children of the current node to the filtered children.
	node.Children = filteredChildren
	// Recursively remove the node from the children nodes.
	for _, child := range node.Children {
		tree.removeNodeRecursive(child, name)
	}
	return node
}

// Traverse traverses the configuration tree using the specified traversal mode and applies the visitor function to each node.
func (tree *ConfigTree) Traverse(mode TreeTraversalMode, visitor ConfigNodeVisitor) error {
	// Switch based on the traversal mode.
	switch mode {
	case PreOrder:
		// Traverse the tree in pre-order.
		return tree.traversePreOrder(tree.Root, visitor)
	case InOrder:
		// Traverse the tree in in-order.
		return tree.traverseInOrder(tree.Root, visitor)
	case PostOrder:
		// Traverse the tree in post-order.
		return tree.traversePostOrder(tree.Root, visitor)
	default:
		// If an unsupported traversal mode is provided, return an error.
		return fmt.Errorf("unsupported traversal mode")
	}
}

// traversePreOrder traverses the configuration tree in pre-order and applies the visitor function to each node.
func (tree *ConfigTree) traversePreOrder(node *ConfigNode, visitor ConfigNodeVisitor) error {
	if node == nil {
		return nil
	}
	// Skip the root node if it has no children
	if node == tree.Root && len(node.Children) == 0 {
		return nil
	}
	// Visit the current node.
	if err := visitor.VisitNode(node); err != nil {
		return err
	}
	// Recursively traverse the children of the current node in pre-order.
	for _, child := range node.Children {
		if err := tree.traversePreOrder(child, visitor); err != nil {
			return err
		}
	}
	return nil
}

// traverseInOrder traverses the configuration tree in in-order and applies the visitor function to each node.
func (tree *ConfigTree) traverseInOrder(node *ConfigNode, visitor ConfigNodeVisitor) error {
	if node == nil {
		return nil
	}
	// Skip the root node if it has no children
	if node == tree.Root && len(node.Children) == 0 {
		return nil
	}
	// Recursively traverse the children of the current node in in-order.
	for _, child := range node.Children {
		if err := tree.traverseInOrder(child, visitor); err != nil {
			return err
		}
	}
	// Visit the current node.
	if err := visitor.VisitNode(node); err != nil {
		return err
	}
	return nil
}

// traversePostOrder traverses the configuration tree in post-order and applies the visitor function to each node.
func (tree *ConfigTree) traversePostOrder(node *ConfigNode, visitor ConfigNodeVisitor) error {
	if node == nil {
		return nil
	}
	// Skip the root node if it has no children
	if node == tree.Root && len(node.Children) == 0 {
		return nil
	}
	// Recursively traverse the children of the current node in post-order.
	for _, child := range node.Children {
		if err := tree.traversePostOrder(child, visitor); err != nil {
			return err
		}
	}
	// Visit the current node.
	if err := visitor.VisitNode(node); err != nil {
		return err
	}
	return nil
}
