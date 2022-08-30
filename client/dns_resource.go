package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

type DNSResource struct {
	ApiClient ApiClient
}

func (d *DNSResource) client(id int) ApiClient {
	return d.ApiClient.GetSubObject(fmt.Sprintf("dnsresources/%v", id))
}

func (d *DNSResource) Get(id int) (dnsResource *entity.DNSResource, err error) {
	dnsResource = new(entity.DNSResource)
	err = d.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})
	return
}

func (d *DNSResource) Update(id int, params *entity.DNSResourceParams) (dnsResource *entity.DNSResource, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	dnsResource = new(entity.DNSResource)
	err = d.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, dnsResource)
	})
	return
}

func (d *DNSResource) Delete(id int) error {
	return d.client(id).Delete()
}
