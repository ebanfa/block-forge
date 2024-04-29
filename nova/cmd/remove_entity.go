/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// removeEntityCmd represents the entity subcommand under the remove command
var removeEntityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Remove an existing entity from the configuration",
	Long:  `Remove an existing entity from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.CommandOptions{
			Debug:   debug,
			Command: plugin.RemoveEntityOp,
		})
	},
}

func init() {
	removeCmd.AddCommand(removeEntityCmd)
}
