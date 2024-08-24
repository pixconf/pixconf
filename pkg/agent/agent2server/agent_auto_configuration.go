package agent2server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pixconf/pixconf/internal/buildinfo"
	"github.com/pixconf/pixconf/pkg/server/proto"
)

func (c *Client) GetAgentAutoConfiguration(ctx context.Context) (*proto.AgentAutoConfigurationResponse, error) {
	relativeURL, err := url.Parse("/.well-known/pixconf/agent-configuration")
	if err != nil {
		return nil, err
	}

	path := c.baseURL.ResolveReference(relativeURL)

	client := &http.Client{}

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, path.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", fmt.Sprintf("pixconf-agent/%s", buildinfo.Version))

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var response = &proto.AgentAutoConfigurationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
