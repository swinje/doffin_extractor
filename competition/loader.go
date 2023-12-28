package competition

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Load() {

	if fileInfo, err := os.Stat("doffin.json"); err != nil {
		refresh()
	} else {
		// file older than 24 hours
		if fileInfo.ModTime().Before(time.Now().Add(-24 * time.Hour)) {
			refresh()
		}

	}

	c, err := LoadCompetition()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.NumHitsTotal)

}

func refresh() {
	fmt.Println("Refreshing data...")
	c, err := GetAPI()
	if err != nil {
		log.Fatal(err)
	}
	StoreCompetition(c)

}
