# Blockchain ETL Package

The Blockchain ETL Package is designed to provide tools for Extract, Transform, and Load (ETL) processes within blockchain applications. It offers functionalities to manage and execute ETL processes efficiently. This package is implemented in Golang and provides a set of interfaces and structures to facilitate the development of ETL pipelines tailored for blockchain data processing.

## Features

- **ETL Process Management**: Manage the lifecycle of ETL processes, including initialization, starting, stopping, restarting, and removal.
- **Configuration Flexibility**: Configurable components and pipelines to adapt to various blockchain data processing requirements.
- **Concurrent Access Handling**: Safely handle concurrent access to ETL processes with mutex locks.
- **Extensible Interfaces**: Interfaces are provided for customization and extension according to specific application needs.

## Components

### `ProcessManagerInterface`
- Interface for managing and executing ETL processes.
- Defines methods for initializing, starting, stopping, restarting, retrieving, and removing ETL processes.

### `ProcessManagerServiceInterface`
- Extends `ProcessManagerInterface` and `SystemServiceInterface`.
- Represents a service for managing ETL processes within the application.

### `ETLProcessConfig`
- Configuration structure for an ETL process, including components configuration.

### `PipelineConfig`
- Configuration structure for a transformation pipeline, including stages configuration.

### `ETLProcessStatus`
- Enum type representing the status of an ETL process.

### `ETLProcess`
- Structure representing an individual ETL process, including its configuration and status.

### `ProcessManagerService`
- Service for managing ETL processes within the application.
- Implements `ProcessManagerInterface` and `SystemServiceInterface`.

### `ProcessManager`
- Implementation of `ProcessManagerService`.
- Manages the lifecycle of ETL processes, handling initialization, starting, stopping, and restarting.

## Usage

### Initialization
```go
// Initialize the ProcessManagerService
processManager := process.NewProcessManagerService("id", "name", "description")

// Initialize ETL processes based on the configuration provided by the system
err := processManager.InitializeProcess(ctx, processConfig)
if err != nil {
    // Handle error
}
```

### Starting ETL Processes
```go
// Start all initialized ETL processes
err := processManager.StartProcess(ctx, processID)
if err != nil {
    // Handle error
}
```

### Stopping ETL Processes
```go
// Stop all running ETL processes
err := processManager.StopProcess(ctx,processID)
if err != nil {
    // Handle error
}
```

### Retrieving ETL Processes
```go
// Get an ETL process by its ID
etlProcess, err := processManager.GetProcess(processID)
if err != nil {
    // Handle error
}
```

### Retrieving ETL All Processes
```go
// Get an ETL process by its ID
etlProcesses, err := processManager.GetAllProcesses()
if err != nil {
    // Handle error
}
```

### Removing ETL Processes
```go
// Remove an ETL process by its ID
err := processManager.RemoveProcess(processID)
if err != nil {
    // Handle error
}
```

---

This README provides an overview of the Blockchain ETL Package, its components, features, and usage guidelines. For detailed documentation and examples, please refer to the source code and documentation comments.

For any questions, issues, or contributions, feel free to contact the package maintainer.

--- 