package competition

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetAPI() (result Competition, err error) {

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{
		Timeout: time.Second * 3,
	}

	var jsonData = []byte(`{
		"numHitsPerPage":2000,
		"page":1,
		"searchString":"",
		"sortBy":"RELEVANCE",
		"facets":{"cpvCode":{"checkedItems":[]},
		"type":{"checkedItems":["COMPETITION"]},
		"status":{"checkedItems":["ACTIVE"]},
		"estimatedValue":{"from":null,"to":null},
		"issueDate":{"from":null,"to":null},
		"location":{"checkedItems":[]}}
		}`)

	doffinURL := "https://dof-search-notices-prod-app.azurewebsites.net/search"
	req, err := http.NewRequest("POST", doffinURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return Competition{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Chrome/51.0.2704.103 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return Competition{}, err
	}

	fmt.Println(resp.Status)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Competition{}, err
	}

	doffin := Competition{}

	err = json.Unmarshal(body, &doffin)
	if err != nil {
		return Competition{}, err
	}

	return doffin, nil
}
