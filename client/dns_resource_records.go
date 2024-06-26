//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// DNSResourceRecords implements api.DNSResourceRecords
type DNSResourceRecords struct {
	APIClient APIClient
}

func (d *DNSResourceRecords) client() APIClient {
	return d.APIClient.GetSubObject("dnsresourcerecords")
}

// Get fetches a list of DNSResourceRecord objectts
func (d *DNSResourceRecords) Get() ([]entity.DNSResourceRecord, error) {
	dnsResourceRecords := make([]entity.DNSResourceRecord, 0)
	err := d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &dnsResourceRecords)
	})

	return dnsResourceRecords, err
}

// Create creates a new DNSResourceRecord
func (d *DNSResourceRecords) Create(params *entity.DNSResourceRecordParams) (*entity.DNSResourceRecord, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	dnsResourceRecord := new(entity.DNSResourceRecord)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResourceRecord)
	})

	return dnsResourceRecord, err
}
