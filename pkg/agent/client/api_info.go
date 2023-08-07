package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pixconf/pixconf/pkg/agent/protos"
)

func (c *Client) GetInfo(ctx context.Context) (*protos.InfoResponse, error) {
	relativeURL, err := url.Parse("/info")
	if err != nil {
		return nil, err
	}

	path := c.baseURL.ResolveReference(relativeURL)

	client := newHTTPClient(c.apiSocketPath)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var infoResponse = &protos.InfoResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&infoResponse); err != nil {
		return nil, err
	}

	return infoResponse, nil
}
