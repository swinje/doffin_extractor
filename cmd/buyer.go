/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/swinje/doffin_extractor/competition"
)

// buyerCmd represents the buyer command
var buyerCmd = &cobra.Command{
	Use:   "buyer",
	Short: "Search by buyer",
	Long:  `Enter buyer name as search term.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := competition.GetCompetition()
		if err != nil {
			return
		}
		if len(args) > 0 {
			r := competition.SearchBuyer(c, args[0])
			fmt.Println("Records matched: ", len(r))
			competition.PrintRecords(c, r)
		} else {
			fmt.Println("No search term.")
		}
	},
}

func init() {
	rootCmd.AddCommand(buyerCmd)
}
