package client

import (
	"encoding/json"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
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
	vmHost = new(entity.VMHost)
	err = m.client().Post("", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})
	return
}
