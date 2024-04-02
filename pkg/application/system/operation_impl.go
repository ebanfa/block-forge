package system

import (
	"errors"
	"fmt"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/context"
)

// SystemOperations is a concrete implementation of the Operations interface.
type SystemOperations struct {
	operations map[string]Operation // Map to store registered operations
	lock       sync.RWMutex         // Read-write mutex to synchronize access to operations
	system     System               // System instance received during initialization
}

// NewSystemOperations creates a new instance of SystemOperations.
func NewSystemOperations() Operations {
	return &SystemOperations{
		operations: make(map[string]Operation),
		lock:       sync.RWMutex{},
	}
}

// ID returns the unique identifier of the component.
func (ops *SystemOperations) ID() string {
	return "system_operations"
}

// Name returns the name of the component.
func (ops *SystemOperations) Name() string {
	return "System Operations"
}

// Description returns the description of the component.
func (ops *SystemOperations) Description() string {
	return "Manages and executes operations within the system"
}

// Initialize initializes the module with the provided system.
func (ops *SystemOperations) Initialize(ctx *context.Context, system System) error {
	ops.system = system
	return nil
}

// RegisterOperation registers an operation with the given ID.
// Returns an error if the operation ID is already registered or if the operation is nil.
func (ops *SystemOperations) RegisterOperation(operationID string, operation Operation) error {
	if operation == nil {
		return errors.New("nil operation provided")
	}

	ops.lock.Lock()
	defer ops.lock.Unlock()

	_, exists := ops.operations[operationID]
	if exists {
		return fmt.Errorf("operation with ID %s already exists", operationID)
	}

	ops.operations[operationID] = operation
	return nil
}

// ExecuteOperation executes the operation with the given ID and input data.
// Returns the output of the operation and an error if the operation is not found or if execution fails.
func (ops *SystemOperations) ExecuteOperation(ctx *context.Context, operationID string, data OperationInput) (OperationOutput, error) {
	ops.lock.RLock()
	operation, exists := ops.operations[operationID]
	ops.lock.RUnlock()
	if !exists {
		return OperationOutput{}, fmt.Errorf("operation with ID %s not found", operationID)
	}

	// Initialize the operation if it implements the Component interface
	err := operation.Initialize(ctx, ops.system)
	if err != nil {
		return OperationOutput{}, fmt.Errorf("failed to initialize operation: %v", err)
	}

	return operation.Execute(ctx, data)
}
