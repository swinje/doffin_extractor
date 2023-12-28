package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/swinje/doffin_extractor/competition"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load data from Doffin",
	Long: `Downloads data from Doffin and stores in local doffin.json file
	It will only download if file is older than 24 hrs`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("load called")
		competition.Load()
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
