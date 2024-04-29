/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Add a new query to the configuration",
	Long:  `Add a new query to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("query called")
	},
}

func init() {
	addCmd.AddCommand(queryCmd)

}
