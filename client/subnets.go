//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Subnets contains functionality for manipulating the Subnets entity.
type Subnets struct {
	APIClient APIClient
}

func (s *Subnets) client() APIClient {
	return s.APIClient.GetSubObject("subnets")
}

// Get fetches a list of Subnet objects
func (s *Subnets) Get() (subnets []entity.Subnet, err error) {
	err = s.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnets)
	})

	return
}

// Create creates a new Subnet
func (s *Subnets) Create(params *entity.SubnetParams) (subnet *entity.Subnet, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	subnet = new(entity.Subnet)
	err = s.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, subnet)
	})

	return
}
