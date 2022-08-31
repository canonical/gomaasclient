package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

type BlockDevices struct {
	ApiClient ApiClient
}

func (b *BlockDevices) client(systemID string) ApiClient {
	return b.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("blockdevices")
}

func (b *BlockDevices) Get(systemID string) (blockDevices []entity.BlockDevice, err error) {
	err = b.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &blockDevices)
	})
	return
}

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
