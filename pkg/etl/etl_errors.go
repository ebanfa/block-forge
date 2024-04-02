package etl

import (
	"errors"
)

var (
	// ErrProcessNotFound is returned when an ETL process is not found.
	ErrProcessNotFound = errors.New("ETL process not found")

	// ErrInvalidProcess is returned when an ETL process is not valid.
	ErrInvalidProcess = errors.New("Invalid ETL Process")

	// ErrScheduledProcessNotFound is returned when a scheduled ETL process is not found.
	ErrScheduledProcessNotFound = errors.New("scheduled ETL process not found")

	ErrETLComponentFactoryExists = errors.New("factory for already exists")

	ErrInvalidETLProcessConfig = errors.New("invalid ETL process config")
)
