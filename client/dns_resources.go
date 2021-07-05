package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type DNSResources struct {
	ApiClient ApiClient
}

func (d *DNSResources) client() ApiClient {
	return d.ApiClient.GetSubObject("dnsresources")
}

func (d *DNSResources) Get() (dnsresources []entity.DNSResource, err error) {
	err = d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &dnsresources)
	})
	return
}

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
