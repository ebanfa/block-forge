/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/spf13/cobra"
)

// createCmd represents the new command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new blockchain project configuration",
	Long: `Create a new blockchain project configuration. This command initializes a 
	new configuration for a blockchain project, allowing you to define and manage the 
	various components and modules of your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.CommandOptions{
			Debug: debug,
		})
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}
