package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// DNSResources implements api.DNSResources
type DNSResources struct {
	ApiClient ApiClient
}

func (d *DNSResources) client() ApiClient {
	return d.ApiClient.GetSubObject("dnsresources")
}

// Get fetches a list of DNSResource objects
func (d *DNSResources) Get() (dnsresources []entity.DNSResource, err error) {
	err = d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &dnsresources)
	})
	return
}

// Create creates a new DNSResource
func (d *DNSResources) Create(params *entity.DNSResourceParams) (dnsResource *entity.DNSResource, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	dnsResource = new(entity.DNSResource)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})
	return
}
