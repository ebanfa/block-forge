# Configuration Repository for Codenet

This repository contains configuration files for Codenet, allowing users to customize and define the structure, entities, messages, and queries of their Cosmos SDK-based blockchains using simple JSON files.

## Table of Contents

1. [Introduction](#introduction)
2. [Directory Structure](#directory-structure)
3. [Configuration Format](#configuration-format)
4. [Usage](#usage)
5. [Contributing](#contributing)
6. [License](#license)

## Introduction

Codenet is a decentralized data encoding network. This repository serves as a centralized location for managing configuration files that define the structure and functionality of Cosmos SDK-based blockchains. By providing a clear structure and format for configuration files, developers can easily customize and deploy their blockchain networks.

## Directory Structure

The repository follows the following directory structure:

```
.
├── config.json                         # Main configuration file for the blockchain project
├── codenet
│   ├── module.config                   # Module-specific configuration file for the codenet module
│   ├── entities                        # Directory for entity configuration files
│   │   ├── encoded-data.json           # Entity configuration for the encoded data
│   │   └── ...
│   ├── messages                        # Directory for message configuration files
│   │   ├── encode-data.json            # Message configuration for encode data transactions
│   │   └── ...
│   └── queries                         # Directory for query configuration files
│       ├── get-encoded-data-by-id.json # Query configuration for retrieving pool information
│       └── ...
└── other_module                        # Directory for configuration files of other modules
    ├── ...
```

## Configuration Format

### Main Configuration File (config.json)

The main configuration file (`config.json`) defines high-level settings and module dependencies for the blockchain project. It includes the following fields:

- `name`: Name of the blockchain project.
- `frontends`: Frontend frameworks used for the project.
- `coins`: Definition of custom coins for the project.
- `abciHandlers`: Definition of ABCI handlers for the project.
- `modules`: List of modules included in the project, with dependencies and directory references for module-specific configuration files.

### Module-Specific Configuration Files

Each module may have its own module-specific configuration file (e.g., `module.config`). These files define module-specific settings and configurations, such as enabling/disabling features or specifying custom parameters.

### Entity Configuration Files

Entity configuration files define the structure of entities within each module. Each entity configuration includes the entity's name, description, and fields with their respective types and descriptions.

### Message Configuration Files

Message configuration files define the structure of messages within each module. Each message configuration includes the message's name, description, and fields with their respective types and descriptions.

### Query Configuration Files

Query configuration files define the structure of queries within each module. Each query configuration includes the query's name, description, and fields with their respective types and descriptions.

## Usage

To use the configuration files in your blockchain project:

1. Clone this repository to your local machine.
2. Customize the configuration files according to your project requirements.
3. Incorporate the configuration files into your blockchain project, ensuring they are referenced correctly.

## Contributing

Contributions to improve the structure and format of the configuration files are welcome! To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your changes.
3. Make your modifications and commit them.
4. Push your changes to your fork.
5. Submit a pull request.

## License

This repository is licensed under the [MIT License](LICENSE).
