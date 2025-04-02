//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// IPRanges implements api.IPRanges
type IPRanges struct {
	APIClient APIClient
}

func (i *IPRanges) client() *APIClient {
	return i.APIClient.SubClient("ipranges")
}

// Get fetches a list of IPRange objects
func (i *IPRanges) Get(ctx context.Context) ([]entity.IPRange, error) {
	ipRanges := make([]entity.IPRange, 0)
	err := i.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &ipRanges)
	})

	return ipRanges, err
}

// Create creates a new IPRange object
func (i *IPRanges) Create(ctx context.Context, params *entity.IPRangeParams) (*entity.IPRange, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	ipRange := new(entity.IPRange)
	err = i.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, ipRange)
	})

	return ipRange, err
}
