//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Subnets contains functionality for manipulating the Subnets entity.
type Subnets struct {
	APIClient APIClient
}

func (s *Subnets) client() APIClient {
	return s.APIClient.GetSubObject("subnets")
}

// Get fetches a list of Subnet objects
func (s *Subnets) Get() ([]entity.Subnet, error) {
	subnets := make([]entity.Subnet, 0)
	err := s.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &subnets)
	})

	return subnets, err
}

// Create creates a new Subnet
func (s *Subnets) Create(params *entity.SubnetParams) (*entity.Subnet, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	subnet := new(entity.Subnet)
	err = s.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, subnet)
	})

	return subnet, err
}
