/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// fieldCmd represents the field command
var fieldCmd = &cobra.Command{
	Use:   "field",
	Short: "Add a new field to the configuration",
	Long:  `Add a new field to the configuration tree`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("field called")
	},
}

func init() {
	addCmd.AddCommand(fieldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fieldCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fieldCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
