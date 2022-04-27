package b2c2

import "encoding/json"

const WITHDRAWALS_ENDPOINT = "withdrawal/"

type DestinationAddress struct {
	AddressValue    string  `json:"address_value"`
	AddressSuffix   *string `json:"address_suffix"`
	AddressProtocol *string `json:"address_protocol"`
}

type withdrawalRequest struct {
	Amount             string              `json:"amount"`
	Currency           string              `json:"currency"`
	DestinationAddress *DestinationAddress `json:"destination_address"`
}

type Withdrawal struct {
	Amount                 string             `json:"amount"`
	Currency               string             `json:"currency"`
	WithdrawalId           string             `json:"withdrawal_id"`
	Reference              string             `json:"reference"`
	Settled                bool               `json:"settled"`
	Created                string             `json:"created"`
	DestinationAddress     DestinationAddress `json:"destination_address"`
	DestinationBankAccount *string            `json:"destination_bank_account"`
}

func (c *Client) Withdrawals() ([]*Withdrawal, error) {
	b, err := c.privateGet(WITHDRAWALS_ENDPOINT)
	if err != nil {
		return nil, err
	}

	result := make([]*Withdrawal, 0)
	err = json.Unmarshal(b, &result)

	return result, err
}

func (c *Client) CreateWithdrawal(currency string, amount string, address string, suffix *string, protocol *string) (*Withdrawal, error) {
	req := withdrawalRequest{
		Amount:   amount,
		Currency: currency,
		DestinationAddress: &DestinationAddress{
			AddressValue:    address,
			AddressSuffix:   suffix,
			AddressProtocol: protocol,
		},
	}

	b, err := c.privatePost(WITHDRAWALS_ENDPOINT, req)
	if err != nil {
		return nil, err
	}

	result := new(Withdrawal)
	err = json.Unmarshal(b, result)

	return result, err
}
