/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Add a new module node to the configuration",
	Long:  `Add a new module node to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("module called")
	},
}

func init() {
	addCmd.AddCommand(moduleCmd)
}
