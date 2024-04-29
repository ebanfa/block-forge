/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize the configuration tree and dependency graph",
	Long:  `Visualize the configuration tree and dependency graph`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.CommandOptions{
			Debug:   debug,
			Command: plugin.VisualizeConfigOp,
		})
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)

}
