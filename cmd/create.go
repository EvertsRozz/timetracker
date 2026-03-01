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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create NAME [WAGE]",
	Short: "Create a new project instance",
	Args:  cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		s, err := store.Load()
		if err != nil {
			return
		}

		projName := strings.TrimSpace(args[0])
		wage, _ := strconv.ParseFloat(args[1], 32)

		proj, err := s.Create(projName, float32(wage))
		if err != nil {
			return
		}
		if err := s.Save(); err != nil {
			return
		}

		fmt.Printf("Project instance named %v saved with wage of %v", proj.Name, proj.Wage)
	},
}

func init() {
	projectCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
