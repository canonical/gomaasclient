package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Domains implements api.Domains
type Domains struct {
	APIClient APIClient
}

func (d *Domains) client() APIClient {
	return d.APIClient.GetSubObject("domains")
}

// Get fetches a list of Domain objects
func (d *Domains) Get() ([]entity.Domain, error) {
	domains := make([]entity.Domain, 0)
	err := d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &domains)
	})

	return domains, err
}

// Create creates a new Domain
func (d *Domains) Create(params *entity.DomainParams) (*entity.Domain, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	domain := new(entity.Domain)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, domain)
	})

	return domain, err
}

// SetSerial sets the SOA serial for all domains
func (d *Domains) SetSerial(serial int) error {
	qsp := url.Values{}
	qsp.Set("serial", fmt.Sprintf("%v", serial))

	return d.client().Post("", qsp, func(data []byte) error { return nil })
}
