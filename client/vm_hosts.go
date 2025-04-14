//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// VMHosts contains functionality for manipulating the VMHosts entity.
type VMHosts struct {
	APIClient APIClient
}

func (p *VMHosts) client() *APIClient {
	return p.APIClient.SubClient("pods")
}

// Get fetches a list of VMHost objects
func (p *VMHosts) Get(ctx context.Context) ([]entity.VMHost, error) {
	vmHosts := make([]entity.VMHost, 0)
	err := p.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &vmHosts)
	})

	return vmHosts, err
}

// Create creates a new VMHost
func (p *VMHosts) Create(ctx context.Context, params *entity.VMHostParams) (*entity.VMHost, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	vmHost := new(entity.VMHost)
	err = p.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})

	return vmHost, err
}
