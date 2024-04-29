/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeEntityCmd represents the entity subcommand under the remove command
var removeEntityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Remove an existing entity from the configuration",
	Long:  `Remove an existing entity from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("entity called under remove command")
	},
}

func init() {
	removeCmd.AddCommand(removeEntityCmd)
}
