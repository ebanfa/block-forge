/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeQueryCmd represents the query subcommand under the remove command
var removeQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Remove an existing query from the configuration",
	Long:  `Remove an existing query from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("query called under remove command")
	},
}

func init() {
	removeCmd.AddCommand(removeQueryCmd)
}
