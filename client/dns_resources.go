//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// DNSResources implements api.DNSResources
type DNSResources struct {
	APIClient APIClient
}

func (d *DNSResources) client() APIClient {
	return d.APIClient.GetSubObject("dnsresources")
}

// Get fetches a list of DNSResource objects
func (d *DNSResources) Get(params *entity.DNSResourcesParams) ([]entity.DNSResource, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	dnsresources := make([]entity.DNSResource, 0)
	err = d.client().Get("", qsp, func(data []byte) error {
		return json.Unmarshal(data, &dnsresources)
	})

	return dnsresources, err
}

// Create creates a new DNSResource
func (d *DNSResources) Create(params *entity.DNSResourceParams) (*entity.DNSResource, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	dnsResource := new(entity.DNSResource)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})

	return dnsResource, err
}
