package competition

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

func GetCompetition() (Competition, error) {
	c, err := LoadCompetition()
	if err != nil {
		log.Println(err)
	}
	return c, err
}

// returns a map of IDs matching
func SearchHeading(c Competition, term string) []int {
	var r []int
	log.Println("Searching for ", term)
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	for i, hit := range c.Hits {
		if re.MatchString(hit.Heading) {
			r = append(r, i)
		}
	}
	return r

}

func SearchDescription(c Competition, term string) []int {
	var r []int
	log.Println("Searching for ", term)
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	for i, hit := range c.Hits {
		desc, _ := hit.Description.(string)
		if re.MatchString(desc) {
			r = append(r, i)
		}

	}
	return r
}

func SearchID(c Competition, term string) []int {
	var r []int
	log.Println("Searching for ", term)
	re := regexp.MustCompile(`(?i)\b` + term + `\b`)
	for i, hit := range c.Hits {
		if re.MatchString(hit.ID) {
			r = append(r, i)
		}
	}
	return r

}

func SearchRecent(c Competition, days string) []int {
	var r []int
	log.Printf("Searching for newer than %s\n", days)

	intDays, err := strconv.Atoi(days)
	if err != nil {
		return r
	}

	for i, hit := range c.Hits {
		if hit.IssueDate.After(time.Now().AddDate(0, 0, -intDays)) {
			r = append(r, i)
		}
	}
	return r
}
