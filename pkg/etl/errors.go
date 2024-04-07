package etl

import "errors"

// Custom errors
var (
	ErrInvalidProcessesConfig = errors.New("invalid processes configuration provided")
	ErrEmptyProcessConfig     = errors.New("process configuration does not contain components")
	ErrNotProcess             = errors.New("provided type is not an ETL Process")
	ErrProcInitOutputInvalid  = errors.New("process initialization output is not an etl process")
	ErrNotProcessComponent    = errors.New("component is not an ETLProcess component")
	ErrNotProcessNotFound     = errors.New("etl process not found")

	// ErrInvalidSource is returned when an invalid or unsupported source is provided.
	ErrInvalidSource = errors.New("invalid or unsupported source")

	// ErrInvalidDestination is returned when an invalid or unsupported destination is provided.
	ErrInvalidDestination = errors.New("invalid or unsupported destination")
)
