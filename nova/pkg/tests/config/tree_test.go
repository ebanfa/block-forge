package config

import (
	"testing"

	"github.com/edward1christian/block-forge/nova/pkg/config"
	"github.com/stretchr/testify/assert"
)

// TestNewConfigTree tests the NewConfigTree function for creating a new configuration tree.
func TestNewConfigTree(t *testing.T) {
	// Act
	tree := config.NewConfigTree()

	// Assert
	assert.NotNil(t, tree)
	assert.NotNil(t, tree.Root)
}

// TestAddNode_AddSingleNode tests adding a single node to the configuration tree.
func TestAddNode_AddSingleNode(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}

	// Act
	tree.AddNode(node)

	// Assert
	assert.NotNil(t, tree.Root)
	assert.NotEqual(t, node, tree.Root)
	assert.Equal(t, 1, len(tree.Root.Children))
}

// TestAddNode_AddMultipleNodes tests adding multiple nodes to the configuration tree.
func TestAddNode_AddMultipleNodes(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node1 := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	node2 := &config.ConfigNode{Name: "node2", Type: config.MessageNode}

	// Act
	tree.AddNode(node1)
	tree.AddNode(node2)

	// Assert
	assert.NotNil(t, tree.Root)
	assert.Equal(t, 2, len(tree.Root.Children))
	assert.Contains(t, tree.Root.Children, node1)
	assert.Contains(t, tree.Root.Children, node2)
}

// TestFindNode_Found tests finding an existing node in the configuration tree.
func TestFindNode_Found(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	tree.AddNode(node)

	// Act
	foundNode := tree.FindNode("node1")

	// Assert
	assert.NotNil(t, foundNode)
	assert.Equal(t, node, foundNode)
}

// TestFindNode_NotFound tests finding a non-existent node in the configuration tree.
func TestFindNode_NotFound(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	tree.AddNode(node)

	// Act
	foundNode := tree.FindNode("nonExistentNode")

	// Assert
	assert.Nil(t, foundNode)
}

// TestRemoveNode_RemoveSingleNode tests removing a single node from the configuration tree.
func TestRemoveNode_RemoveSingleNode(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	tree.AddNode(node)

	// Act
	tree.RemoveNode("node1")

	// Assert
	assert.NotNil(t, tree.Root)
	assert.Equal(t, 0, len(tree.Root.Children))
}

// TestRemoveNode_RemoveMultipleNodes tests removing multiple nodes from the configuration tree.
func TestRemoveNode_RemoveMultipleNodes(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node1 := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	node2 := &config.ConfigNode{Name: "node2", Type: config.MessageNode}

	tree.AddNode(node1)
	tree.AddNode(node2)

	// Act
	tree.RemoveNode("node1")

	// Assert
	assert.NotNil(t, tree.Root)
	assert.Equal(t, 1, len(tree.Root.Children))
	assert.NotContains(t, tree.Root.Children, node1)
	assert.Contains(t, tree.Root.Children, node2)
}

// MockVisitor is a mock implementation of the ConfigNodeVisitor interface for testing purposes.
type MockVisitor struct {
	VisitedNodes []*config.ConfigNode
}

// VisitNode mocks visiting a node and stores the visited node.
func (v *MockVisitor) VisitNode(node *config.ConfigNode) error {
	v.VisitedNodes = append(v.VisitedNodes, node)
	return nil
}

// TestTraverse_PreOrder tests traversing the configuration tree in pre-order.
func TestTraverse_PreOrder(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node1 := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	node2 := &config.ConfigNode{Name: "node2", Type: config.MessageNode}

	tree.AddNode(node1)
	tree.AddNode(node2)
	visitor := &MockVisitor{}

	// Act
	tree.Traverse(config.PreOrder, visitor)

	// Assert
	assert.NotNil(t, visitor.VisitedNodes)
	assert.Equal(t, 3, len(visitor.VisitedNodes))
	/* assert.Equal(t, node1, visitor.VisitedNodes[1])
	assert.Equal(t, node2, visitor.VisitedNodes[2]) */
}

// TestTraverse_InOrder tests traversing the configuration tree in in-order.
func TestTraverse_InOrder(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node1 := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	node2 := &config.ConfigNode{Name: "node2", Type: config.MessageNode}

	tree.AddNode(node1)
	tree.AddNode(node2)
	visitor := &MockVisitor{}

	// Act
	tree.Traverse(config.InOrder, visitor)

	// Assert
	assert.NotNil(t, visitor.VisitedNodes)
	assert.Equal(t, 3, len(visitor.VisitedNodes))
	/* assert.Equal(t, node1, visitor.VisitedNodes[1])
	assert.Equal(t, node2, visitor.VisitedNodes[2]) */
}

// TestTraverse_PostOrder tests traversing the configuration tree in post-order.
func TestTraverse_PostOrder(t *testing.T) {
	// Arrange
	tree := config.NewConfigTree()
	node1 := &config.ConfigNode{Name: "node1", Type: config.ModuleNode}
	node2 := &config.ConfigNode{Name: "node2", Type: config.MessageNode}

	tree.AddNode(node1)
	tree.AddNode(node2)
	visitor := &MockVisitor{}

	// Act
	tree.Traverse(config.PostOrder, visitor)

	// Assert
	assert.NotNil(t, visitor.VisitedNodes)
	assert.Equal(t, 3, len(visitor.VisitedNodes))
	/* assert.Equal(t, node1, visitor.VisitedNodes[1])
	assert.Equal(t, node2, visitor.VisitedNodes[2]) */
}
