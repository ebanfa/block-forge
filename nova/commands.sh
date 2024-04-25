#!/bin/bash

# Navigate to your project directory where your main.go file is located
cd /path/to/your/project

# Define the path to the cobra-cli command
COBRA_CLI="$HOME/go/bin/cobra-cli"

# Add top-level commands
$COBRA_CLI add init
$COBRA_CLI add generate
$COBRA_CLI add build
$COBRA_CLI add run
$COBRA_CLI add config

# Add subcommands for config
$COBRA_CLI add new -p configCmd
$COBRA_CLI add load -p configCmd
$COBRA_CLI add save -p configCmd
$COBRA_CLI add add -p configCmd
$COBRA_CLI add update -p configCmd
$COBRA_CLI add remove -p configCmd
$COBRA_CLI add validate -p configCmd
$COBRA_CLI add visualize -p configCmd
$COBRA_CLI add dependency -p configCmd

# Add subcommands for config add
$COBRA_CLI add module -p addCmd
$COBRA_CLI add message -p addCmd
$COBRA_CLI add type -p addCmd
$COBRA_CLI add query -p addCmd
$COBRA_CLI add field -p addCmd

# Add subcommands for config dependency
$COBRA_CLI add link -p dependencyCmd
$COBRA_CLI add unlink -p dependencyCmd
