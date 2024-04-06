package cmd

import (
	"fmt"

	"github.com/edward1christian/block-forge/pkg/application"
	"github.com/spf13/cobra"
)

var (
	appConfigFile       string
	frameworkConfigFile string
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		//internal.Init()
		application.Init(appConfigFile, frameworkConfigFile)
	},
}

func init() {
	startCmd.Flags().StringVar(&frameworkConfigFile, "framework-config", "f", "Path to the framework configuration file (optional)")
	startCmd.Flags().StringVar(&appConfigFile, "app-config", "a", "Path to the application configuration file")

	// Mark the "app-config" flag as required
	startCmd.MarkFlagRequired("app-config")
	//startCmd.MarkFlagRequired("framework-config")

	rootCmd.AddCommand(startCmd)
}
