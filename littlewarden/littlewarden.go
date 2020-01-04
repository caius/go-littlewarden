package littlewarden

import (
	"net/url"
)

const (
	defaultApiURL = "https://littlewarden.com/api"
)

// Main client everything else hangs off
type Client struct {
	ApiKey  string
	BaseURL *url.URL
}

// Build a new client with default attributes set
func NewClient(apikey string) (*Client, error) {
	baseURL, err := url.Parse(defaultApiURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		ApiKey:  apikey,
		BaseURL: baseURL,
	}

	return c, nil
}
