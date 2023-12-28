package competition

import "time"

type Competition struct {
	NumHitsTotal      int `json:"numHitsTotal"`
	NumHitsAccessible int `json:"numHitsAccessible"`
	Hits              []struct {
		ID    string `json:"id"`
		Buyer []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"buyer"`
		Heading          string    `json:"heading"`
		Description      any       `json:"description"`
		LocationID       []any     `json:"locationId"`
		EstimatedValue   any       `json:"estimatedValue"`
		Type             string    `json:"type"`
		AllTypes         []string  `json:"allTypes"`
		Status           string    `json:"status"`
		IssueDate        time.Time `json:"issueDate"`
		Deadline         time.Time `json:"deadline"`
		ReceivedTenders  any       `json:"receivedTenders"`
		LimitedDataFlag  bool      `json:"limitedDataFlag"`
		DoffinClassicURL any       `json:"doffinClassicUrl"`
	} `json:"hits"`
	Facets struct {
		Status struct {
			Items []struct {
				Value string `json:"value"`
				Total int    `json:"total"`
			} `json:"items"`
			TotalRemaining int `json:"totalRemaining"`
		} `json:"status"`
		Type struct {
			Items []struct {
				Value string `json:"value"`
				Total int    `json:"total"`
			} `json:"items"`
			TotalRemaining int `json:"totalRemaining"`
		} `json:"type"`
		Locations struct {
			Items          []any `json:"items"`
			TotalRemaining int   `json:"totalRemaining"`
		} `json:"locations"`
		CpvCode struct {
			Items []struct {
				Value string `json:"value"`
				Total int    `json:"total"`
			} `json:"items"`
			TotalRemaining int `json:"totalRemaining"`
		} `json:"cpvCode"`
	} `json:"facets"`
}
