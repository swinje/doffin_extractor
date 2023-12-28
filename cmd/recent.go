/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/swinje/doffin_extractor/competition"
)

// recentCmd represents the recent command
var recentCmd = &cobra.Command{
	Use:   "recent",
	Short: "Search for recent postings",
	Long:  `Enter days back as term`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := competition.GetCompetition()
		if err != nil {
			return
		}
		if len(args) > 0 {
			r := competition.SearchRecent(c, args[0])
			fmt.Println("Records matched: ", len(r))
			competition.PrintRecords(c, r)
		} else {
			fmt.Println("No search term.")
		}
	},
}

func init() {
	rootCmd.AddCommand(recentCmd)
}
