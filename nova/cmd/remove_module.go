/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// removeModuleCmd represents the module subcommand under the remove command
var removeModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Remove an existing module node from the configuration",
	Long:  `Remove an existing module node from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.CommandOptions{
			Debug:   debug,
			Command: plugin.RemoveModuleOp,
		})
	},
}

func init() {
	removeCmd.AddCommand(removeModuleCmd)
}
