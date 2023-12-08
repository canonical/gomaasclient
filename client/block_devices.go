package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// BlockDevices implements api.BlockDevices
type BlockDevices struct {
	APIClient APIClient
}

func (b *BlockDevices) client(systemID string) APIClient {
	return b.APIClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("blockdevices")
}

// Get fetches a list of BlockDevices for a given system_id
func (b *BlockDevices) Get(systemID string) (blockDevices []entity.BlockDevice, err error) {
	err = b.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &blockDevices)
	})

	return
}

// Create creates a new BlockDevice for a given system_id
func (b *BlockDevices) Create(systemID string, params *entity.BlockDeviceParams) (blockDevice *entity.BlockDevice, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID).Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return
}
