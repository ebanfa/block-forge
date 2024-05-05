/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Add a new query to the configuration",
	Long:  `Add a new query to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.InitOptions{
			Debug:   debug,
			Command: plugin.AddQueryOp,
		})
	},
}

func init() {
	addCmd.AddCommand(queryCmd)

}
