/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log PROJECT MINUTES [NOTE]",
	Short: "Add a log entry to a project",
	Long:  `Add a log entry to a project. Automatically adds the current time to the log.`,
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		projName := args[0]
		minutes64, err := strconv.ParseUint(args[1], 10, 16)
		if err != nil {
			fmt.Printf("❌ Invalid minutes '%s': %v\n", args[1], err)
			return
		}
		minutes := uint16(minutes64)

	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// logCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// logCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
