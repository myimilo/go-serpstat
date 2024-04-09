package serpstat

import "net/http"

type Client struct {
	config ClientConfig

	HTTPClient *http.Client
}

func NewClient(authToken string) *Client {
	config := DefaultConfig(authToken)
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:     config,
		HTTPClient: &http.Client{},
	}
}
