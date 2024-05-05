/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// messageCmd represents the message command
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Add a new message to the configuration",
	Long:  `Add a new message to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.InitOptions{
			Debug:   debug,
			Command: plugin.AddMessageOp,
		})
	},
}

func init() {
	addCmd.AddCommand(messageCmd)

}
