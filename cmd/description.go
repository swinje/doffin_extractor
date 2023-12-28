package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/swinje/doffin_extractor/competition"
)

// descriptionCmd represents the description command
var descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "Search by description",
	Long:  `Enter description as search term.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := competition.GetCompetition()
		if err != nil {
			return
		}
		if len(args) > 0 {
			r := competition.SearchDescription(c, args[0])
			fmt.Println("Records matched: ", len(r))
			competition.PrintRecords(c, r)
		} else {
			fmt.Println("No search term.")
		}
	},
}

func init() {
	rootCmd.AddCommand(descriptionCmd)
}
