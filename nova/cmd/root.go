/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	force      bool
	debug      bool
	verbose    bool
	configFile string
	outputDir  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "buildnet",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
			examples and usage of using your application. For example:

			Cobra is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		//internal.Init()
		//application.Init(appConfigFile)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.buildnet.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Specify the path to the configuration file")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Specify the output directory for generated files")
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite of existing files in the output directory")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Enable verbose mode for detailed output")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "Enable debug mode for troubleshooting")

	// Add flags for help and version
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of Codenet")

}
