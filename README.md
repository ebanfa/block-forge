# BlockForge

BlockForge is a monolithic repository (monorepo) for developing and maintaining a suite of blockchain-related applications and libraries. This monorepo serves as a centralized hub for various blockchain projects, enabling efficient collaboration, code reuse, and streamlined development workflows.

## Table of Contents

- [Overview](#overview)
- [Projects](#projects)
- [Getting Started](#getting-started)
- [Contributing](#contributing)
- [License](#license)

## Overview

BlockForge is designed to foster the development of cutting-edge blockchain solutions, ranging from data relays and blockchain building tools to decentralized bridges, business process management systems, and ZK Rollup chains. By consolidating multiple projects into a single repository, we aim to promote code sharing, consistent dependency management, and a unified development experience.

## Projects

The following projects are currently hosted within the BlockForge monorepo:

- **Necta**: An  ETL (Extract, Transform, Load) tool for blockchain data. Necta can play a crucial role in ensuring data availability and accessibility when dealing with parent and child blockchains
- **Nova**: A blockchain building tool that enables rapid scaffolding of Cosmos SDK-based blockchains.

Additional projects, such as a decentralized bridge, a decentralized business process management blockchain, and a ZK Rollup chain, are planned for future development within the monorepo.

## Getting Started

To get started with BlockForge, follow these steps:

1. **Prerequisites**: Ensure you have Go installed on your machine. You can download it from the official [Go website](https://golang.org/dl/).

2. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/blockforge.git
   ```

3. **Navigate to the monorepo directory**:

   ```bash
   cd blockforge
   ```

4. **Install dependencies**:

   ```bash
   go mod download
   ```

5. **Build and run a project**:

   ```bash
   # Build and run Necta
   go build -o necta ./necta/cmd
   ./necta

   # Build and run Nova
   go build -o nova ./nova/cmd
   ./nova
   ```

For more detailed instructions on building, testing, and running specific projects, refer to the respective project's documentation within the monorepo.

## Contributing

We welcome contributions to the BlockForge monorepo! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix: `git checkout -b my-feature-branch`.
3. Make your changes and commit them: `git commit -am 'Add some feature'`.
4. Push your changes to your fork: `git push origin my-feature-branch`.
5. Create a new Pull Request on the main repository.

Please make sure to follow our [Code of Conduct](CODE_OF_CONDUCT.md) and [Contributing Guidelines](CONTRIBUTING.md) when contributing to this project.

## License

BlockForge is licensed under the [MIT License](LICENSE).

