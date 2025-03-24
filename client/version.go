package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// Version implements api.Version
type Version struct {
	APIClient APIClient
}

func (v *Version) client() *APIClient {
	return v.APIClient.SubClient("version")
}

// Get fetches MAAS version details
func (v *Version) Get(ctx context.Context) (*entity.Version, error) {
	version := new(entity.Version)
	err := v.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, version)
	})

	return version, err
}
