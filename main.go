package main

import (
	"github.com/swinje/doffin_extractor/cmd"
	"github.com/swinje/doffin_extractor/competition"
)

func main() {
	_, err := competition.GetCompetition()
	if err != nil {
		competition.Load()
	}
	cmd.Execute()

}
