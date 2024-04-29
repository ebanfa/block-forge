/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeMessageCmd represents the message subcommand under the remove command
var removeMessageCmd = &cobra.Command{
	Use:   "message",
	Short: "Remove an existing message from the configuration",
	Long:  `Remove an existing message from the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("message called under remove command")
	},
}

func init() {
	removeCmd.AddCommand(removeMessageCmd)
}
