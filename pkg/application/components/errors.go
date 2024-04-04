package components

import "errors"

// Custom errors
var (
	ErrComponentNil                  = errors.New("component is nil")
	ErrComponentAlreadyExist         = errors.New("component already exists")
	ErrFactoryNotFound               = errors.New("factory not found")
	ErrComponentFactoryNil           = errors.New("component factory is nil")
	ErrComponentFactoryAlreadyExists = errors.New("component factory already exists")
	ErrComponentNotFound             = errors.New("component not found")
	ErrComponentTypeNotFound         = errors.New("component type not found")
)
