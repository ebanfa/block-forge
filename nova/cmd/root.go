/*
Copyright © 2024 Edward Banfa <ebanfa@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	debug   bool
	verbose bool
	daemon  bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nova",
	Short: "A Modular Blockchain Framework Scaffolding and Configuration Management Application",
	Long: `Nova is an application designed to streamline the development of modular blockchain applications. It provides a comprehensive set of tools and features for scaffolding, configuring, and generating code for blockchain frameworks that follow a modular architecture.

Nova empowers developers to define and manage the configuration of various components and modules that make up their blockchain application, such as transactions, queries, state structures, and dependencies. With Nova, you can:

- Define and configure the core components of your blockchain application through a user-friendly command-line interface or graphical user interface.
- Leverage a modular and composable approach to building blockchain applications, promoting code reusability and maintainability.
- Validate and visualize your application's configuration, ensuring consistency and correctness before code generation.
- Generate boilerplate code, data structures, and artifacts specific to your target blockchain platform or framework, automating the scaffolding process.
- Support for multiple blockchain platforms and frameworks, including Cosmos SDK, Ethereum (as smart contracts), Substrate, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
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
	// Define flags for command-line options
	rootCmd.PersistentFlags().BoolVar(&daemon, "daemon", false, "Run in daemon mode")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug mode for troubleshooting")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Enable verbose mode for detailed output")

	// Add flags for help and version
	rootCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version of Codenet")

	// Mark config and output flags as required
	//_ = rootCmd.MarkFlagRequired("config")

}
