package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type BlockDevicePartitions struct {
	ApiClient ApiClient
}

func (p *BlockDevicePartitions) client(systemID string, blockDeviceID int) ApiClient {
	return p.ApiClient.GetSubObject(fmt.Sprintf("nodes/%s/blockdevices/%v/partitions", systemID, blockDeviceID))
}

func (p *BlockDevicePartitions) Get(systemID string, blockDeviceID int) (partitions []entity.BlockDevicePartition, err error) {
	err = p.client(systemID, blockDeviceID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &partitions)
	})
	return
}

func (p *BlockDevicePartitions) Create(systemID string, blockDeviceID int, params *entity.BlockDevicePartitionParams) (partition *entity.BlockDevicePartition, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	partition = new(entity.BlockDevicePartition)
	err = p.client(systemID, blockDeviceID).Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}
