package b2c2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_URL = "https://api.b2c2.net"
const SANDBOX_URL = "https://api.uat.b2c2.net"

const AUTHORIZATION = "Authorization"
const TOKEN = "Token %s"

type ApiError struct {
	Code int
	Body string
}

func newApiError(code int, message []byte) *ApiError {
	return &ApiError{
		Code: code,
		Body: string(message),
	}
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("API Error: Code(%d) %v", e.Code, e.Body)
}

type Client struct {
	apiToken string
	apiUrl   string
}

func NewClient(token string) *Client {
	return &Client{apiToken: token, apiUrl: API_URL}
}

func (c *Client) Sandbox() {
	c.apiUrl = SANDBOX_URL
}

func (c *Client) addAuthorization(req *http.Request) {
	req.Header.Add(AUTHORIZATION, fmt.Sprintf(TOKEN, c.apiToken))

	if req.Method == http.MethodPost {
		req.Header.Add("content-type", "application/json")
	}
}

func (c *Client) privateGet(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%v/%v", c.apiUrl, endpoint)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	c.addAuthorization(req)
	return c.doRequest(req)
}

func (c *Client) privateGetWithParams(endpoint string, parameters map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%v/%v", c.apiUrl, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range parameters {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	c.addAuthorization(req)
	return c.doRequest(req)
}

func (c *Client) privatePost(endpoint string, data any) ([]byte, error) {
	url := fmt.Sprintf("%v/%v", c.apiUrl, endpoint)
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(b))
	buf := bytes.NewReader(b)

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, err
	}

	c.addAuthorization(req)
	return c.doRequest(req)
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 300 {
		apiErr := newApiError(resp.StatusCode, b)
		return nil, apiErr
	}

	return b, nil
}
