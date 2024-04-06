package process

import (
	"fmt"
	"sync"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/components"
	systemApi "github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl"
	"github.com/edward1christian/block-forge/pkg/etl/common"
)

// ProcessManager implements the ETLManagerService interface.
type ProcessManager struct {
	systemApi.BaseSystemComponent
	mutex     sync.Mutex             // Mutex for concurrent access to maps
	processes map[string]*ETLProcess // Map to store ETL processes by ID
}

// NewETLManagerService creates a new instance of ETLManagerService.
func NewETLManagerService(id, name, description string) *ProcessManager {
	return &ProcessManager{
		processes: make(map[string]*ETLProcess),
		BaseSystemComponent: systemApi.BaseSystemComponent{
			BaseComponent: components.BaseComponent{
				Id:   id,
				Nm:   name,
				Desc: description,
			},
		},
	}
}

// InitializeProcess initializes the ETLManagerService.
func (m *ProcessManager) InitializeProcess(ctx *context.Context, config *ETLProcessConfig) (*ETLProcess, error) {
	system := m.BaseSystemComponent.System

	// Execute the operation
	output, err := system.ExecuteOperation(ctx, common.ProcessOpInitializeETL, &systemApi.OperationInput{
		Data: config,
	})
	if err != nil {
		return nil, fmt.Errorf("error executing system operation: %v", err)
	}

	// Ensure the process is of type ETLProcess
	etlProcess, ok := output.Data.(*ETLProcess)
	if !ok {
		return nil, etl.ErrNotProcessComponent
	}

	// Add the process to the map
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.processes[etlProcess.ID] = etlProcess

	return etlProcess, nil
}

// StartProcess starts an ETL process with the given ID.
func (m *ProcessManager) StartProcess(ctx *context.Context, processID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, found := m.processes[processID]
	if !found {
		return fmt.Errorf("etl process: %s not found", processID)
	}

	system := m.BaseSystemComponent.System

	// Execute the operation
	_, err := system.ExecuteOperation(ctx, common.ProcessOpStartETL, &systemApi.OperationInput{
		Data: processID,
	})
	if err != nil {
		return fmt.Errorf("error starting process: %v", err)
	}

	return nil
}

// StopProcess stops an ETL process with the given ID.
func (m *ProcessManager) StopProcess(ctx *context.Context, processID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, found := m.processes[processID]
	if !found {
		return fmt.Errorf("etl process: %s not found", processID)
	}

	system := m.BaseSystemComponent.System

	// Execute the operation
	_, err := system.ExecuteOperation(ctx, common.ProcessOpStopETL, &systemApi.OperationInput{
		Data: processID,
	})
	if err != nil {
		return fmt.Errorf("error starting process: %v", err)
	}
	return nil
}

// RestartProcess restarts an ETL process with the given ID.
func (m *ProcessManager) RestartProcess(ctx *context.Context, processID string) error {
	_, found := m.processes[processID]
	if !found {
		return etl.ErrNotProcessNotFound
	}

	// Stop the ETL process
	if err := m.StopProcess(ctx, processID); err != nil {
		return err
	}

	// Start the ETL process
	if err := m.StartProcess(ctx, processID); err != nil {
		return err
	}

	return nil
}

// GetProcess retrieves an ETL process by its ID.
func (m *ProcessManager) GetProcess(processID string) (*ETLProcess, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	process, found := m.processes[processID]
	if !found {
		return nil, etl.ErrNotProcessNotFound
	}
	return process, nil
}

// GetAllProcesses retrieves all ETL processes.
func (m *ProcessManager) GetAllProcesses() []*ETLProcess {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	processes := make([]*ETLProcess, 0, len(m.processes))
	for _, process := range m.processes {
		processes = append(processes, process)
	}
	return processes
}

// RemoveProcess removes an ETL process with the given ID
func (m *ProcessManager) RemoveProcess(processID string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, found := m.processes[processID]; !found {
		return etl.ErrNotProcessNotFound
	}
	delete(m.processes, processID)
	return nil
}
