package competition

import (
	"encoding/json"
	"os"
)

func StoreCompetition(c Competition) (err error) {

	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile("doffin.json", file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadCompetition() (c Competition, err error) {

	data, err := os.ReadFile("doffin.json")
	if err != nil {
		return Competition{}, err
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		return Competition{}, err
	}

	return c, nil
}
