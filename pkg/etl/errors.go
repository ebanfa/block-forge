package etl

import "errors"

// Custom errors
var (
	ErrInvalidProcessesConfig = errors.New("invalid process configuration provided")
	ErrEmptyProcessConfig     = errors.New("process configuration does not contain components")
	ErrNotProcess             = errors.New("provided type is not an ETL Process")
	ErrProcInitOutputInvalid  = errors.New("process initialization output is not an etl process")
	ErrNotProcessComponent    = errors.New("component is not an ETLProcess component")
	ErrNotProcessNotFound     = errors.New("etl process not found")
)
