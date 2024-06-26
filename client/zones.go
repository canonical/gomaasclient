//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Zones implements api.Zones
type Zones struct {
	APIClient APIClient
}

func (z *Zones) client() APIClient {
	return z.APIClient.GetSubObject("zones")
}

// Get fetches a list of Zone objects
func (z *Zones) Get() ([]entity.Zone, error) {
	zones := make([]entity.Zone, 0)
	err := z.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &zones)
	})

	return zones, err
}

// Create creates a new Zone
func (z *Zones) Create(params *entity.ZoneParams) (*entity.Zone, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	zone := new(entity.Zone)
	err = z.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, zone)
	})

	return zone, err
}
