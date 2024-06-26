//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// DNSResource implements api.DNSResource
type DNSResource struct {
	APIClient APIClient
}

func (d *DNSResource) client(id int) APIClient {
	return d.APIClient.GetSubObject(fmt.Sprintf("dnsresources/%v", id))
}

// Get fetches a given DNSResource
func (d *DNSResource) Get(id int) (*entity.DNSResource, error) {
	dnsResource := new(entity.DNSResource)
	err := d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})

	return dnsResource, err
}

// Update updates a given DNSResource
func (d *DNSResource) Update(id int, params *entity.DNSResourceParams) (*entity.DNSResource, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	dnsResource := new(entity.DNSResource)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})

	return dnsResource, err
}

// Delete deletes a given DNSResource
func (d *DNSResource) Delete(id int) error {
	return d.client(id).Delete()
}
