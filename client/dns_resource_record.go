//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// DNSResourceRecord implements api.DNSResourceRecord
type DNSResourceRecord struct {
	APIClient APIClient
}

func (d *DNSResourceRecord) client(id int) APIClient {
	return d.APIClient.GetSubObject(fmt.Sprintf("dnsresourcerecords/%v", id))
}

// Get fetches a given DNSResourceRecord
func (d *DNSResourceRecord) Get(id int) (*entity.DNSResourceRecord, error) {
	dnsResourceRecord := new(entity.DNSResourceRecord)
	err := d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, dnsResourceRecord)
	})

	return dnsResourceRecord, err
}

// Update updates a given DNSResourceRecord
func (d *DNSResourceRecord) Update(id int, params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	dnsResourceRecord := new(entity.DNSResourceRecord)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResourceRecord)
	})

	return dnsResourceRecord, err
}

// Delete deletes a given DNSResourceRecord
func (d *DNSResourceRecord) Delete(id int) error {
	return d.client(id).Delete()
}
