package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// DNSResourceRecord implements api.DNSResourceRecord
type DNSResourceRecord struct {
	ApiClient ApiClient
}

func (d *DNSResourceRecord) client(id int) ApiClient {
	return d.ApiClient.GetSubObject(fmt.Sprintf("dnsresourcerecords/%v", id))
}

// Get fetches a given DNSResourceRecord
func (d *DNSResourceRecord) Get(id int) (dnsResourceRecord *entity.DNSResourceRecord, err error) {
	dnsResourceRecord = new(entity.DNSResourceRecord)
	err = d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, dnsResourceRecord)
	})

	return
}

// Update updates a given DNSResourceRecord
func (d *DNSResourceRecord) Update(id int, params *entity.DNSResourceRecordParams) (dnsResourceRecord *entity.DNSResourceRecord, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	dnsResourceRecord = new(entity.DNSResourceRecord)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResourceRecord)
	})

	return
}

// Delete deletes a given DNSResourceRecord
func (d *DNSResourceRecord) Delete(id int) error {
	return d.client(id).Delete()
}
