package client

import "net/url"

type Client struct {
	apiSocketPath string
	baseURL       *url.URL
}

type Options struct {
	APISocketPath string
}

func NewClient(opts Options) (*Client, error) {
	baseURL, err := url.Parse("http://pixconf-agent")
	if err != nil {
		return nil, err
	}

	return &Client{
		apiSocketPath: opts.APISocketPath,
		baseURL:       baseURL,
	}, nil
}
