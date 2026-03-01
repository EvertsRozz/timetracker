/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/EvertsRozz/timetracker/internal/store"
	"github.com/spf13/cobra"
)

var createFlag bool

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log PROJECT MINUTES [NOTE]",
	Short: "Add a log entry to a project",
	Long:  `Add a log entry to a project. Automatically adds the current time to the log.`,
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.Load()
		if err != nil {
			return
		}

		projectName := strings.TrimSpace(args[0])
		minutes64, err := strconv.ParseUint(args[1], 10, 16)
		if err != nil {
			return
		}
		minutes := uint16(minutes64)

		note := ""
		if len(args) > 2 {
			note = strings.TrimSpace(args[2])
		}

		// Find existing
		proj, err := s.Find(projectName, createFlag) // Pass flag
		if err != nil {
			fmt.Printf("Project %v not found \n", projectName)
			return
		}

		proj.AddLog(minutes, note)
		if err := s.Save(); err != nil {
			return
		}

		fmt.Printf("Logged %d min to %s (Total: %d min)\n",
			minutes, proj.Name, proj.TotalTime)
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
