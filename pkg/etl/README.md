# Blockchain ETL Package

This GoLang package provides functionality for managing and executing Extract, Transform, Load (ETL) processes in the context of blockchain data. It offers an interface for initializing, starting, stopping, restarting, and managing ETL processes.

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [API Reference](#api-reference)
- [Contributing](#contributing)
- [License](#license)

## Overview

The package includes several key components:

- **Process Management**: Interfaces and implementations for managing ETL processes.
- **ETL Process**: Structs and methods for representing individual ETL processes, including initialization, status tracking, and configuration.
- **Scheduled ETL Process**: Structs and methods for scheduling ETL processes for execution at specific intervals.
- **Process Manager Service**: Service for managing ETL processes, including initialization, starting, and stopping of processes.
- **Pipeline Configuration**: Configuration for defining transformation pipelines.

## Installation

To install the package, you can use Go modules:

```bash
go get github.com/edward1christian/block-forge/pkg/etl/process
```

## Usage

Here's a basic example demonstrating how to use the package:

```go
package main

import (
	"fmt"
	"time"

	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl/process"
)

func main() {
	// Initialize context and system
	ctx := context.NewContext()
	sys := system.NewSystem()

	// Initialize process manager service
	manager := process.NewProcessManagerService("1", "Process Manager", "Manages ETL processes")
	pms := services.NewProcessManagerService("1", "Process Manager", "Manages ETL processes", manager)

	// Initialize and start ETL processes
	if err := pms.Initialize(ctx, sys); err != nil {
		fmt.Println("Failed to initialize ETL processes:", err)
		return
	}
	if err := pms.Start(ctx); err != nil {
		fmt.Println("Failed to start ETL processes:", err)
		return
	}

	// Stop ETL processes after a duration
	time.Sleep(10 * time.Second)
	if err := pms.Stop(ctx); err != nil {
		fmt.Println("Failed to stop ETL processes:", err)
		return
	}
}
```

## API Reference

### Interfaces

- **`ProcessManagerInterface`**: Interface for managing and executing ETL processes.

### Structs

- **`ETLProcessConfig`**: Configuration for an ETL process, including component configurations.
- **`PipelineConfig`**: Configuration for a transformation pipeline.
- **`ETLProcessStatus`**: Enum representing the status of an ETL process.
- **`ETLProcess`**: Struct representing an individual ETL process, including its ID, status, configuration, and instantiated components.

### Services

- **`ProcessManagerService`**: Service for managing ETL processes, implementing.
- **`ProcessManager`**: Implementation of `ProcessManagerInterface`, providing methods for initializing, starting, stopping, and managing ETL processes.

## Contributing

Contributions to this package are welcome! Feel free to submit issues or pull requests.

## License

This package is licensed under the [MIT License](LICENSE).

---