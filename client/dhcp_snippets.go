//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// DHCPSnippets implements api.DHCPSnippets
type DHCPSnippets struct {
	APIClient APIClient
}

func (d *DHCPSnippets) client() APIClient {
	return d.APIClient.GetSubObject("dhcp-snippets")
}

// Get fetches a list of DHCPSnippet objects
func (d *DHCPSnippets) Get() ([]entity.DHCPSnippet, error) {
	dhcpSnippets := make([]entity.DHCPSnippet, 0)
	err := d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &dhcpSnippets)
	})

	return dhcpSnippets, err
}

// Create creates a new DHCPSnippet
func (d *DHCPSnippets) Create(params *entity.DHCPSnippetParams) (*entity.DHCPSnippet, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	snippet := new(entity.DHCPSnippet)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, snippet)
	})

	return snippet, err
}
