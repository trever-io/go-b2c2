package b2c2

import "encoding/json"

const BALANCE_ENDPOINT = "balance/"

type Balance map[string]string

func (c *Client) Balance() (Balance, error) {
	b, err := c.privateGet(BALANCE_ENDPOINT)
	if err != nil {
		return nil, err
	}

	balance := Balance{}
	err = json.Unmarshal(b, &balance)
	return balance, err
}
