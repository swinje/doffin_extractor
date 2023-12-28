package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/swinje/doffin_extractor/competition"
)

// idCmd represents the id command
var idCmd = &cobra.Command{
	Use:   "id",
	Short: "Search by ID",
	Long:  `Enter ID as search term.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := competition.GetCompetition()
		if err != nil {
			return
		}
		if len(args) > 0 {
			r := competition.SearchID(c, args[0])
			fmt.Println("Records matched: ", len(r))
			competition.PrintRecords(c, r)
		} else {
			fmt.Println("No search term.")
		}
	},
}

func init() {
	rootCmd.AddCommand(idCmd)
}
