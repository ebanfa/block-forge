Nova implements the following commands, subcommands, and flags:

1. **Top-Level Commands**:
   - `init`: Initialize a new blockchain project.
   - `generate`: Generate code and artifacts for the blockchain application based on the defined configuration.
   - `build`: Build the blockchain application binary.
   - `run`: Run the blockchain application.
   - `config`: Manage the configuration of the blockchain application.

2. **Subcommands for `config`**:
   - `config new`: Create a new configuration tree.
   - `config load`: Load an existing configuration tree.
   - `config save`: Save the current configuration tree.
   - `config add`: Add a new node to the configuration tree.
   - `config update`: Update an existing node in the configuration tree.
   - `config remove`: Remove a node from the configuration tree.
   - `config validate`: Validate the current configuration tree.
   - `config visualize`: Visualize the configuration tree and dependency graph.
   - `config dependency`: Manage dependencies between configuration nodes.

3. **Subcommands for `config add`**:
   - `config add module`: Add a new module node to the configuration tree.
   - `config add transaction`: Add a new transaction node to the configuration tree.
   - `config add query`: Add a new query node to the configuration tree.
   - `config add field`: Add a new field node to the configuration tree.

4. **Subcommands for `config dependency`**:
   - `config dependency add`: Add a dependency between two configuration nodes.
   - `config dependency remove`: Remove a dependency between two configuration nodes.

5. **Flags**:
   - `--name`: Specify the name of the node being added or updated.
   - `--type`: Specify the type of the node being added (module, transaction, query, field).
   - `--parent`: Specify the parent node for the new node being added.
   - `--input-fields`: Specify the input fields for a transaction or query node (comma-separated list).
   - `--output-fields`: Specify the output fields for a transaction node (comma-separated list).
   - `--response-field`: Specify the response field for a query node.
   - `--handler`: Specify the handler function for a transaction node.
   - `--dependency`: Specify the dependency node when adding or removing a dependency.
   - `--config-file`: Specify the path to the configuration file to load or save.
   - `--format`: Specify the format for visualizing the configuration tree and dependency graph (e.g., ASCII, GraphViz).
   - `--verbose`: Enable verbose output for debugging and logging.

Here's an example of how these commands and subcommands could be used:

```
# Initialize a new blockchain project
nova-cli init my-project

# Add a new module
nova-cli config add module --name=token --parent=root

# Add a new transaction to the token module
nova-cli config add transaction --name=transfer --parent=token --input-fields=sender,recipient,amount --handler=handleTransfer

# Add a new query to the token module
nova-cli config add query --name=balance --parent=token --input-fields=address --response-field=balance

# Add a dependency between the token module and the governance module
nova-cli config dependency add --dependency=governance --parent=token

# Validate the current configuration
nova-cli config validate

# Visualize the configuration tree and dependency graph
nova-cli config visualize --format=ascii

# Generate code and artifacts for the blockchain application
nova-cli generate

# Build the blockchain application binary
nova-cli build

# Run the blockchain application
nova-cli run
```
