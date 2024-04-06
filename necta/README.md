# Necta: Blockchain ETL Tool

Necta is a versatile blockchain ETL (Extract, Transform, Load) tool integrated into the Block Forge monorepo. It facilitates the transfer of data between different blockchain networks by extracting data from source blockchains, performing necessary transformations, and loading the data into target blockchains. Necta is built using Go and leverages an event-based architecture, plugin-based architecture, and orchestration framework for flexibility, extensibility, and manageability.

## Table of Contents

- [Overview](#overview)
- [Components](#components)
  - [Blockchain Adapters](#blockchain-adapters)
  - [Transformation Pipeline](#transformation-pipeline)
  - [Blockchain Relays](#blockchain-relays)
  - [Event Bus](#event-bus)
  - [Plugin System](#plugin-system)
  - [Orchestration Framework](#orchestration-framework)
  - [Task Scheduler](#task-scheduler)
  - [Monitoring and Logging](#monitoring-and-logging)
- [Key Features](#key-features)
- [Use Cases](#use-cases)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running Necta](#running-necta)
- [Contributing](#contributing)
- [License](#license)
- [Conclusion](#conclusion)

## Overview

Necta, serves as a powerful blockchain ETL tool within the Block Forge monorepo. It facilitates the extraction of data from source blockchain networks, performs necessary transformations, and loads the data into target blockchain networks. Necta is designed to handle diverse blockchain data transfer needs efficiently.

## Components

### Blockchain Adapters

Necta includes blockchain-specific adapters responsible for extracting data from source blockchain networks. These adapters, communicate with the source blockchain's APIs to fetch data such as transactions, blocks, and smart contract events. Each adapter publishes events to the event bus to notify other components of new data availability.

### Transformation Pipeline

Extracted data flows through a transformation pipeline consisting of multiple stages. The pipeline, composed of modular components implemented as plugins, allows flexible and customizable data transformations. Each pipeline stage subscribes to relevant events on the event bus and processes data asynchronously.

### Blockchain Relays

Processed data is forwarded to blockchain-specific relays responsible for loading the data into the target blockchain networks. Relays, implemented as plugins, interact with the target blockchain's APIs to submit transactions, deploy smart contracts, or update on-chain data. Each relay subscribes to events on the event bus to receive data for loading into the target blockchain.

### Event Bus

Necta utilizes an event bus to facilitate communication between loosely coupled components. Components publish events to the event bus to notify others of changes or new data availability. Subscribers listen for specific events and react accordingly, enabling decoupled and asynchronous communication.

### Plugin System

Necta features a plugin-based architecture allowing dynamic loading and unloading of functionality. New adapters, pipelines, relays, and event handlers can be added via plugins without modifying the core codebase. Plugins are dynamically discovered and loaded at runtime, providing flexibility and extensibility for customization.

### Orchestration Framework

Necta incorporates an orchestration framework responsible for coordinating and executing ETL processes. The framework interprets ETL configurations defined by users and coordinates the execution of adapters, pipelines, and relays accordingly. It handles task scheduling, dependency resolution, error handling, and monitoring of ETL processes.

### Task Scheduler

Necta integrates a task scheduler component to enable the scheduling of ETL processes at predefined intervals or specified triggers. The task scheduler invokes the orchestration framework to execute scheduled tasks based on defined schedules and configurations.

### Monitoring and Logging

The system includes monitoring and logging capabilities to track the progress and performance of ETL processes. Monitoring tools provide real-time insights into the status of adapters, pipelines, and relays, allowing proactive management and troubleshooting.

## Key Features

- **Event-Based Architecture**: Necta employs an event-based architecture for asynchronous and decoupled data processing.
- **Plugin-Based Architecture**: The plugin-based architecture allows dynamic addition of new functionality without modifying the core codebase.
- **Orchestration Framework**: Necta provides a formal mechanism for orchestrating and managing ETL processes, enabling task scheduling, dependency resolution, and monitoring.
- **Modular Design**: Necta is designed with modularity in mind, promoting code reuse, maintainability, and scalability.
- **Scalability**: Necta enables horizontal scalability to handle increasing volumes of data and workload efficiently.

## Use Cases

- **Blockchain Migration**: Necta facilitates seamless data migration between different blockchain networks.
- **Cross-Chain Integration**: The tool synchronizes data between multiple blockchain networks, customizable via plugins.
- **Blockchain Analytics**: Necta enables blockchain analytics and data analysis by extracting, transforming, and loading blockchain data into analytics platforms or databases.
- **Blockchain Development**: Developers leverage Necta for blockchain development workflows, such as testing and deploying smart contracts across different blockchain networks.

## Getting Started

This section provides instructions for setting up and running Necta on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) version 1.16 or later
- [Git](https://git-scm.com/downloads) for cloning the repository

### Installation

1. Clone the Block Forge monorepo:

   ```bash
   git clone https://github.com/blockforge/monorepo.git
   ```

2. Navigate to the `necta` directory:

   ```bash
   cd monorepo/necta
   ```

3. Build the Necta binary:

   ```bash
   go build
   ```

### Configuration

Necta requires a configuration file to specify the desired adapters, pipelines, relays, and other settings. Refer to the [Configuration Guide](docs/configuration.md) for detailed instructions on setting up your configuration.

### Running Necta

After building the binary and configuring Necta, you can run it with the following command:

```bash
./necta --config /path/to/your/config.yaml
```

Replace `/path/to/your/config.yaml` with the actual path to your configuration file.

## Contributing

We welcome contributions to Necta! If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on the [Block Forge monorepo](https://github.com/blockforge/monorepo). For detailed contribution guidelines, refer to the [Contributing Guide](CONTRIBUTING.md).

## License

Necta is licensed under the [MIT License](LICENSE).

## Conclusion

Necta, formerly BitRelay, is a comprehensive blockchain ETL tool integrated into the Block Forge monorepo. With its event-based and plugin-based architecture, complemented by an orchestration framework, Necta streamlines blockchain data transfer and integration effectively. It empowers users to manage and execute ETL processes with flexibility, extensibility, and manageability, enhancing blockchain workflows within the Block Forge ecosystem.