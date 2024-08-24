package agent2server

import "net/url"

type Client struct {
	baseURL *url.URL
}

type Options struct {
	ServerEndpoint string
}

func NewClient(opts Options) (*Client, error) {
	baseURL, err := url.Parse(opts.ServerEndpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseURL: baseURL,
	}, nil
}
