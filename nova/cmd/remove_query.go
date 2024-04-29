/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// removeQueryCmd represents the query subcommand under the remove command
var removeQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Remove an existing query from the configuration",
	Long:  `Remove an existing query from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.CommandOptions{
			Debug:   debug,
			Command: plugin.RemoveQueryOp,
		})
	},
}

func init() {
	removeCmd.AddCommand(removeQueryCmd)
}
