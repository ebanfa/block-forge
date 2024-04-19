package system

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/component"
)

// BaseSystemService represents a concrete implementation of the OperationInterface.
type BaseSystemService struct {
	BaseSystemComponent
}

// Type returns the type of the component.
func (bo *BaseSystemService) Type() component.ComponentType {
	return component.ServiceType
}

// NewBaseSystemService creates a new instance of BaseSystemService.
func NewBaseSystemService(id, name, description string) *BaseSystemService {
	return &BaseSystemService{
		BaseSystemComponent: BaseSystemComponent{
			BaseComponent: component.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// Start starts the component.
// Returns an error if the start operation fails.
func (bo *BaseSystemService) Start(ctx *context.Context) error {
	// Start the service component
	return errors.New("service not implemented")
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (bo *BaseSystemService) Stop(ctx *context.Context) error {
	// Stop the service component
	return errors.New("service not implemented")
}
