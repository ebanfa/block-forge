/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// removeMessageCmd represents the message subcommand under the remove command
var removeMessageCmd = &cobra.Command{
	Use:   "message",
	Short: "Remove an existing message from the configuration",
	Long:  `Remove an existing message from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.InitOptions{
			Debug:   debug,
			Command: plugin.RemoveMessageOp,
		})
	},
}

func init() {
	removeCmd.AddCommand(removeMessageCmd)
}
