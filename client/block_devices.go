//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BlockDevices implements api.BlockDevices
type BlockDevices struct {
	APIClient APIClient
}

func (b *BlockDevices) client(systemID string) *APIClient {
	return b.APIClient.SubClient("nodes").SubClient(systemID).SubClient("blockdevices")
}

// Get fetches a list of BlockDevices for a given system_id
func (b *BlockDevices) Get(ctx context.Context, systemID string) ([]entity.BlockDevice, error) {
	blockDevices := make([]entity.BlockDevice, 0)
	err := b.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &blockDevices)
	})

	return blockDevices, err
}

// Create creates a new BlockDevice for a given system_id
func (b *BlockDevices) Create(ctx context.Context, systemID string, params *entity.BlockDeviceParams) (*entity.BlockDevice, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	blockDevice := new(entity.BlockDevice)
	err = b.client(systemID).Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}
