package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
)

// BlockDevicePartition implements the api.BlockDevicePartition
// interface
type BlockDevicePartition struct {
	APIClient APIClient
}

func (p *BlockDevicePartition) client(systemID string, blockDeviceID int, id int) *APIClient {
	return p.APIClient.
		SubClient("nodes").SubClient(systemID).
		SubClient("blockdevices").SubClient(fmt.Sprintf("%v", blockDeviceID)).
		SubClient("partition").SubClientWithoutSlash(fmt.Sprintf("%v", id))
}

// Get fetches a given BlockDevicePartion for a given system_id, BlockDevice id and partition id
func (p *BlockDevicePartition) Get(ctx context.Context, systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error) {
	partition := new(entity.BlockDevicePartition)
	err := p.client(systemID, blockDeviceID, id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}

// Delete deletes a given BlockDevicePartition
func (p *BlockDevicePartition) Delete(ctx context.Context, systemID string, blockDeviceID int, id int) error {
	return p.client(systemID, blockDeviceID, id).Delete(ctx)
}

// AddTag adds a tag to a given BlockDevicePartition
func (p *BlockDevicePartition) AddTag(ctx context.Context, systemID string, blockDeviceID int, id int, tag string) (*entity.BlockDevicePartition, error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)

	partition := new(entity.BlockDevicePartition)
	err := p.client(systemID, blockDeviceID, id).Post(ctx, "add_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}

// RemoveTag removes a tag from a given BlockDevicePartition
func (p *BlockDevicePartition) RemoveTag(ctx context.Context, systemID string, blockDeviceID int, id int, tag string) (*entity.BlockDevicePartition, error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)

	partition := new(entity.BlockDevicePartition)
	err := p.client(systemID, blockDeviceID, id).Post(ctx, "remove_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}

// Format sets the BlockDevicePartition to be formatted with a given file system
func (p *BlockDevicePartition) Format(ctx context.Context, systemID string, blockDeviceID int, id int, fsType string, label string) (*entity.BlockDevicePartition, error) {
	qsp := url.Values{}
	qsp.Set("fstype", fsType)
	qsp.Set("label", label)

	partition := new(entity.BlockDevicePartition)
	err := p.client(systemID, blockDeviceID, id).Post(ctx, "format", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}

// Unformat unsets a file system for a given BlockDevicePartition
func (p *BlockDevicePartition) Unformat(ctx context.Context, systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error) {
	partition := new(entity.BlockDevicePartition)
	err := p.client(systemID, blockDeviceID, id).Post(ctx, "unformat", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}

// Mount sets a mount point and options for a given BlockDevicePartition
func (p *BlockDevicePartition) Mount(ctx context.Context, systemID string, blockDeviceID int, id int, mountPoint string, mountOptions string) (*entity.BlockDevicePartition, error) {
	qsp := url.Values{}
	qsp.Set("mount_point", mountPoint)
	qsp.Set("mount_options", mountOptions)

	partition := new(entity.BlockDevicePartition)
	err := p.client(systemID, blockDeviceID, id).Post(ctx, "mount", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}

// Unmount unsets a mount point for a given BlockDevicePartition
func (p *BlockDevicePartition) Unmount(ctx context.Context, systemID string, blockDeviceID int, id int) (*entity.BlockDevicePartition, error) {
	partition := new(entity.BlockDevicePartition)
	err := p.client(systemID, blockDeviceID, id).Post(ctx, "unmount", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})

	return partition, err
}
