/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeModuleCmd represents the module subcommand under the remove command
var removeModuleCmd = &cobra.Command{
	Use:   "module",
	Short: "Remove an existing module node from the configuration",
	Long:  `Remove an existing module node from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("module called under remove command")
	},
}

func init() {
	removeCmd.AddCommand(removeModuleCmd)
}
