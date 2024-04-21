/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	provider "github.com/edward1christian/block-forge/nova/pkg"
	"github.com/spf13/cobra"
)

var (
	force          bool
	debug          bool
	verbose        bool
	configFilePath string
	outputDir      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nova",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		provider.Init(&provider.InitOptions{
			Debug:          debug,
			Verbose:        verbose,
			ConfigFilePath: configFilePath,
		})
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nova.yaml)")

	// Define flags for command-line options
	rootCmd.Flags().StringVarP(&configFilePath, "config", "c", "", "Specify the path to the configuration file")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Specify the output directory for generated files")
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite of existing files in the output directory")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose mode for detailed output")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode for troubleshooting")

	// Add flags for help and version
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of Codenet")

	// Mark config and output flags as required
	_ = rootCmd.MarkFlagRequired("config")
	_ = rootCmd.MarkFlagRequired("output")
}
