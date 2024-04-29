/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Add a new message to the configuration",
	Long:  `Add a new message to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("message called")
	},
}

func init() {
	addCmd.AddCommand(messageCmd)

}
