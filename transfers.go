package b2c2

import (
	"encoding/json"
)

const LEDGER_ENDPOINT = "ledger/"

type Transfer struct {
	Transaction_id string `json:"transaction_id"`
	Created        string `json:"created"`
	Reference      string `json:"reference"`
	Currency       string `json:"currency"`
	Amount         string `json:"amount"`
	Type           string `json:"type"`
	Group          string `json:"group"`
}

func (c *Client) Transfers(startTime string, endTime string) ([]*Transfer, error) {
	parameters := map[string]string{
		"created__gte": startTime,
		"created__lt":  endTime,
		"type":         "transfer",
		"limit":        "1000",
	}

	b, err := c.privateGetWithParams(LEDGER_ENDPOINT, parameters)
	if err != nil {
		return nil, err
	}

	result := make([]*Transfer, 0)
	err = json.Unmarshal(b, &result)

	return result, err
}
