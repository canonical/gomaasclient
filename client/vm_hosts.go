package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// VMHosts contains functionality for manipulating the VMHosts entity.
type VMHosts struct {
	APIClient APIClient
}

func (p *VMHosts) client() APIClient {
	return p.APIClient.GetSubObject("pods")
}

// Get fetches a list of VMHost objects
func (p *VMHosts) Get() (vmHosts []entity.VMHost, err error) {
	err = p.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &vmHosts)
	})

	return
}

// Create creates a new VMHost
func (p *VMHosts) Create(params *entity.VMHostParams) (vmHost *entity.VMHost, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	vmHost = new(entity.VMHost)
	err = p.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, vmHost)
	})

	return
}
