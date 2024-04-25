/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
