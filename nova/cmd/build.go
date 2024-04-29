/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the blockchain application binary",
	Long:  `Build the blockchain application binary`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

}
