//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// DHCPSnippet implements api.DHCPSnippet
type DHCPSnippet struct {
	APIClient APIClient
}

func (d *DHCPSnippet) client(id int) APIClient {
	return d.APIClient.GetSubObject(fmt.Sprintf("dhcp-snippets/%v", id))
}

// Get fetches a given DHCPSnippet
func (d *DHCPSnippet) Get(id int) (*entity.DHCPSnippet, error) {
	snippet := new(entity.DHCPSnippet)
	err := d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, snippet)
	})

	return snippet, err
}

// Update updates the given DHCPSnippet
func (d *DHCPSnippet) Update(id int, params *entity.DHCPSnippetParams) (*entity.DHCPSnippet, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	snippet := new(entity.DHCPSnippet)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, snippet)
	})

	return snippet, err
}

// Delete deletes a given DHCPSnippet
func (d *DHCPSnippet) Delete(id int) error {
	return d.client(id).Delete()
}
