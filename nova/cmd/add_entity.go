/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// entityCmd represents the type command
var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Add a new entity to the configuration",
	Long:  `Add a new entity to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("entity called")
	},
}

func init() {
	addCmd.AddCommand(entityCmd)

}
