//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BootSource implements api.BootSource
type BootSource struct {
	APIClient APIClient
}

func (b *BootSource) client(id int) APIClient {
	return b.APIClient.GetSubObject("boot-sources").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a boot source with a given id
func (b *BootSource) Get(id int) (*entity.BootSource, error) {
	bootSource := new(entity.BootSource)
	err := b.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, bootSource)
	})

	return bootSource, err
}

// Update updates a given boot source
func (b *BootSource) Update(id int, params *entity.BootSourceParams) (*entity.BootSource, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	bootSource := new(entity.BootSource)
	err = b.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, bootSource)
	})

	return bootSource, err
}

// Delete deletes a given boot source
func (b *BootSource) Delete(id int) error {
	return b.client(id).Delete()
}
