package competition

import (
	"github.com/fatih/color"
)

func PrintRecords(c Competition, idx []int) {

	for i := range idx {
		j := idx[i]
		color.White("\n------------------------------")
		color.Red("Anbud ID: %s", c.Hits[j].ID)
		color.Blue("Tittel: %s", c.Hits[j].Heading)
		if c.Hits[j].Description != nil {
			desc, _ := c.Hits[j].Description.(string)
			color.Yellow(desc)
		}
		color.Cyan("Publisert: %s", c.Hits[j].IssueDate.Format("02.01 2006"))
		color.Cyan("Frist: %s", c.Hits[j].Deadline.Format("02.01 2006"))
		for _, b := range c.Hits[j].Buyer {
			color.Magenta("Kjøper: %s", b.Name)
		}
		color.Cyan("Status: %s", c.Hits[j].Status)
		if c.Hits[j].EstimatedValue != nil {
			ev := c.Hits[j].EstimatedValue.(map[string]interface{})
			color.Cyan("Verdi: %.0f %s", ev["amount"], ev["currencyCode"])
		}
		color.Cyan("URL: https://doffin.no/notices/%s", c.Hits[j].ID)

	}
}
