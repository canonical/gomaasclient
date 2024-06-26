//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BootSources implements api.BootSources
type BootSources struct {
	APIClient APIClient
}

func (b *BootSources) client() APIClient {
	return b.APIClient.GetSubObject("boot-sources")
}

// Get fetches a list of boot sources
func (b *BootSources) Get() ([]entity.BootSource, error) {
	bootSources := make([]entity.BootSource, 0)
	err := b.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &bootSources)
	})

	return bootSources, err
}

// Create creates a new boot source
func (b *BootSources) Create(params *entity.BootSourceParams) (*entity.BootSource, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bootSource := new(entity.BootSource)
	err = b.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, bootSource)
	})

	return bootSource, err
}
