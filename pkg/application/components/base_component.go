package components

// BaseComponent represents a concrete implementation of the OperationInterface.
type BaseComponent struct {
	id          string
	name        string
	description string
}

func NewComponentImpl(id, name, description string) *BaseComponent {
	return &BaseComponent{id: id, name: name, description: description}
}

// ID returns the unique identifier of the component.
func (oc *BaseComponent) ID() string {
	return oc.id
}

// Name returns the name of the component.
func (oc *BaseComponent) Name() string {
	return oc.name
}

// Type returns the type of the component.
func (oc *BaseComponent) Type() ComponentType {
	return BasicComponentType
}

// Description returns the description of the component.
func (oc *BaseComponent) Description() string {
	return oc.description
}
