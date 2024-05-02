/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/edward1christian/block-forge/nova/pkg/utils"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/spf13/cobra"
)

// createCmd represents the new command
var createCmd = &cobra.Command{
	Use:   "create [projectName]",
	Short: "Create a new blockchain project configuration",
	Long: `Create a new blockchain project configuration. This command initializes a 
    new configuration for a blockchain project, allowing you to define and manage the 
    various components and modules of your application.`,
	Args: cobra.ExactArgs(1), // Example expects exactly 1 argument
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Validate input
		if projectName == "" {
			fmt.Println("Project name is required.")
			return
		}

		// Get the home directory of the current user
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Failed to get user's home directory: %v\n", err)
			return
		}

		// Create SystemOperationInput with project ID and home directory
		inputData := &system.SystemOperationInput{
			Data: map[string]interface{}{
				"projectID":   utils.HashSHA256(projectName),
				"homeDir":     homeDir, // Pass the home directory here
				"projectName": projectName,
			},
		}

		// Populate CommandOptions with arguments and input data
		commandOptions := provider.CommandOptions{
			Debug:   debug,
			Data:    inputData,
			Command: plugin.CreateConfigurationOp,
		}

		// Pass CommandOptions to your main application API
		provider.Init(&commandOptions)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
