/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configurations",
	Long:  `List all configurations`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.InitOptions{
			Debug:    debug,
			Command:  plugin.ListConfigurationsOp,
			InitMode: provider.CommandMode,
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}
