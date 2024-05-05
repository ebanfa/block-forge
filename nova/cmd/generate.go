/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/edward1christian/block-forge/nova/pkg/components/plugin"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code and artifacts for the blockchain application based on the defined configuration",
	Long:  `Generate code and artifacts for the blockchain application based on the defined configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		provider.Init(&provider.InitOptions{
			Debug:   debug,
			Command: plugin.GenerateArtifactsOp,
		})
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
