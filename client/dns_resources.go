//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/canonical/gomaasclient/entity"
)

// DNSResources implements api.DNSResources
type DNSResources struct {
	APIClient APIClient
}

func (d *DNSResources) client() APIClient {
	return d.APIClient.GetSubObject("dnsresources")
}

// Get fetches a list of DNSResource objects
func (d *DNSResources) Get() ([]entity.DNSResource, error) {
	dnsresources := make([]entity.DNSResource, 0)
	err := d.client().Get("", url.Values{}, func(data []byte) error {
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
