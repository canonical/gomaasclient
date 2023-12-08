//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// DNSResourceRecords implements api.DNSResourceRecords
type DNSResourceRecords struct {
	APIClient APIClient
}

func (d *DNSResourceRecords) client() APIClient {
	return d.APIClient.GetSubObject("dnsresourcerecords")
}

// Get fetches a list of DNSResourceRecord objectts
func (d *DNSResourceRecords) Get() (dnsResourceRecords []entity.DNSResourceRecord, err error) {
	err = d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &dnsResourceRecords)
	})

	return
}

// Create creates a new DNSResourceRecord
func (d *DNSResourceRecords) Create(params *entity.DNSResourceRecordParams) (dnsResourceRecord *entity.DNSResourceRecord, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	dnsResourceRecord = new(entity.DNSResourceRecord)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResourceRecord)
	})

	return
}
