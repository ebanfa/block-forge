/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize the configuration tree and dependency graph",
	Long:  `Visualize the configuration tree and dependency graph`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("visualize called")
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)

}
