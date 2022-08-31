package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Contains functionality for manipulating the VMHosts entity.
type VMHosts struct {
	ApiClient ApiClient
}

func (p *VMHosts) client() ApiClient {
	return p.ApiClient.GetSubObject("pods")
}

func (p *VMHosts) Get() (vmHosts []entity.VMHost, err error) {
	err = p.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &vmHosts)
	})
	return
}

func (m *VMHosts) Create(params *entity.VMHostParams) (vmHost *entity.VMHost, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	vmHost = new(entity.VMHost)
	err = m.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})
	return
}
