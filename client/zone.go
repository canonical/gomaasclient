package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Zone Contains functionality for manipulating the Zone entity.
type Zone struct {
	APIClient APIClient
}

func (z *Zone) client(name string) *APIClient {
	return z.APIClient.SubClient("zones").SubClient(name)
}

// Get Zone details.
func (z *Zone) Get(ctx context.Context, name string) (*entity.Zone, error) {
	zone := new(entity.Zone)
	err := z.client(name).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, zone)
	})

	return zone, err
}

// Update Zone.
func (z *Zone) Update(ctx context.Context, name string, params *entity.ZoneParams) (*entity.Zone, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	zone := new(entity.Zone)
	err = z.client(name).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, zone)
	})

	return zone, err
}

// Delete Zone.
func (z *Zone) Delete(ctx context.Context, name string) error {
	return z.client(name).Delete(ctx)
}
