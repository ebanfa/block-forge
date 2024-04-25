#!/bin/bash

# Navigate to your project directory where your main.go file is located
cd /path/to/your/project

# Add top-level commands
cobra-cli add init
cobra-cli add generate
cobra-cli add build
cobra-cli add run
cobra-cli add config

# Add subcommands for config
cobra-cli add new -p configCmd
cobra-cli add load -p configCmd
cobra-cli add save -p configCmd
cobra-cli add add -p configCmd
cobra-cli add update -p configCmd
cobra-cli add remove -p configCmd
cobra-cli add validate -p configCmd
cobra-cli add visualize -p configCmd
cobra-cli add dependency -p configCmd

# Add subcommands for config add
cd add
cobra-cli add module
cobra-cli add transaction
cobra-cli add query
cobra-cli add field

# Add subcommands for config dependency
cd ../dependency
cobra-cli add add
cobra-cli add remove

# Navigate back to the project root directory
cd ../../..
