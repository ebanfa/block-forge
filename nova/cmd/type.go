/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// typeCmd represents the type command
var typeCmd = &cobra.Command{
	Use:   "type",
	Short: "Add a new entity to the configuration",
	Long:  `Add a new entity to the configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("type called")
	},
}

func init() {
	addCmd.AddCommand(typeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// typeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// typeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
