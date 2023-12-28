package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/swinje/doffin_extractor/competition"
)

// headingCmd represents the heading command
var headingCmd = &cobra.Command{
	Use:   "heading",
	Short: "Search by heading",
	Long:  `Enter heading as search term.`,
	Run: func(cmd *cobra.Command, args []string) {
		c, err := competition.GetCompetition()
		if err != nil {
			return
		}
		if len(args) > 0 {
			r := competition.SearchHeading(c, args[0])
			fmt.Println("Records matched: ", len(r))
			competition.PrintRecords(c, r)
		} else {
			fmt.Println("No search term.")
		}

	},
}

func init() {
	rootCmd.AddCommand(headingCmd)
}
