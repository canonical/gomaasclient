package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

type DNSResourceRecords struct {
	ApiClient ApiClient
}

func (d *DNSResourceRecords) client() ApiClient {
	return d.ApiClient.GetSubObject("dnsresourcerecords")
}

func (d *DNSResourceRecords) Get() (dnsResourceRecords []entity.DNSResourceRecord, err error) {
	err = d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &dnsResourceRecords)
	})
	return
}

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
