package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Zone Contains functionality for manipulating the Zone entity.
type Zone struct {
	APIClient APIClient
}

func (z *Zone) client(name string) APIClient {
	return z.APIClient.GetSubObject("zones").GetSubObject(name)
}

// Get Zone details.
func (z *Zone) Get(name string) (*entity.Zone, error) {
	zone := new(entity.Zone)
	err := z.client(name).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, zone)
	})

	return zone, err
}

// Update Zone.
func (z *Zone) Update(name string, params *entity.ZoneParams) (*entity.Zone, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	zone := new(entity.Zone)
	err = z.client(name).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, zone)
	})

	return zone, err
}

// Delete Zone.
func (z *Zone) Delete(name string) error {
	return z.client(name).Delete()
}
