package services

import (
	"errors"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	"github.com/edward1christian/block-forge/pkg/application/system"
)

// BaseComponent represents a concrete implementation of the OperationInterface.
type DemoService struct {
	system.BaseSystemService // Embedding BaseComponent
}

// Type returns the type of the component.
func (bo *DemoService) Type() components.ComponentType {
	return components.ServiceType
}

func NewDemoService(id, name, description string) *DemoService {
	return &DemoService{
		BaseSystemService: system.BaseSystemService{
			BaseSystemComponent: system.BaseSystemComponent{
				BaseComponent: components.BaseComponent{
					Id:   id,
					Nm:   name,
					Desc: description,
				},
			},
		},
	}
}

// Start starts the component.
// Returns an error if the start operation fails.
func (bo *DemoService) Start(ctx *context.Context) error {
	// Start the service component
	return errors.New("service not implemented")
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (bo *DemoService) Stop(ctx *context.Context) error {
	// Stop the service component
	return errors.New("service not implemented")
}
