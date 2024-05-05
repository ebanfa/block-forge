/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Add a new module node to the configuration",
	Long:  `Add a new module node to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.InitOptions{
			Debug:   debug,
			Command: plugin.AddModuleOp,
		})
	},
}

func init() {
	addCmd.AddCommand(moduleCmd)
}
