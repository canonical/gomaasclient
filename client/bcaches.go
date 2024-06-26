//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/canonical/gomaasclient/entity"
)

// BCaches contains functionality for manipulating the BCaches entity.
type BCaches struct {
	APIClient APIClient
}

func (b *BCaches) client(systemID string) APIClient {
	return b.APIClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("bcaches")
}

// Get BCaches of a machine.
func (b *BCaches) Get(systemID string) ([]entity.BCache, error) {
	bCaches := make([]entity.BCache, 0)
	err := b.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &bCaches)
	})

	return bCaches, err
}

// Create a BCache of a machine.
func (b *BCaches) Create(systemID string, params *entity.BCacheParams) (*entity.BCache, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bCache := new(entity.BCache)
	err = b.client(systemID).Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, bCache)
	})

	return bCache, err
}
