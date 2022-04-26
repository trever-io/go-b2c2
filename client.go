package b2c2

type Client struct {
	apiToken string
}

func NewClient(token string) *Client {
	return &Client{apiToken: token}
}
