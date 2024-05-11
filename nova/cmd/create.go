/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"errors"

	"github.com/asaskevich/govalidator"
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
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
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get project name
		projectName := args[0]

		// Validate input
		if !govalidator.IsAlphanumeric(projectName) {
			return errors.New("project name must be alphanumeric")
		}

		// Pass InitOptions to your main application API
		provider.Init(&provider.InitOptions{
			Debug:   debug,
			Daemon:  daemon,
			Verbose: verbose,
			Command: plugin.CreateConfigurationOp,
			Data:    projectName,
		})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}
