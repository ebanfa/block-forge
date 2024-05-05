/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// entityCmd represents the type command
var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Add a new entity to the configuration",
	Long:  `Add a new entity to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.InitOptions{
			Debug:   debug,
			Command: plugin.RemoveEntityOp,
		})
	},
}

func init() {
	addCmd.AddCommand(entityCmd)

}
