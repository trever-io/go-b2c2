package b2c2

import (
	"encoding/json"
)

const CURRENCIES_ENDPOINT = "currency/"

type CurrencyInfo struct {
	StableCoin       bool    `json:"stable_coin"`
	IsCrypto         bool    `json:"is_crypto"`
	CurrencyType     string  `json:"currency_type"`
	ReadableName     string  `json:"readable_name"`
	LongOnly         bool    `json:"long_only"`
	MinimumTradeSide float64 `json:"minimum_trade_side"`
}

type Currencies map[string]CurrencyInfo

func (c *Client) Currencies() (Currencies, error) {
	b, err := c.privateGet(CURRENCIES_ENDPOINT)
	if err != nil {
		return nil, err
	}

	currencies := Currencies{}
	err = json.Unmarshal(b, &currencies)
	return currencies, err
}
