package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// BlockDevice implements the api.BlockDevice interface
type BlockDevice struct {
	APIClient APIClient
}

func (b *BlockDevice) client(systemID string, id int) *APIClient {
	return b.APIClient.SubClient("nodes").SubClient(systemID).SubClient("blockdevices").SubClient(fmt.Sprintf("%v", id))
}

// Get fetches a given BlockDevice based on the given Node's system_id and BlockDevice's id
func (b *BlockDevice) Get(ctx context.Context, systemID string, id int) (*entity.BlockDevice, error) {
	blockDevice := new(entity.BlockDevice)
	err := b.client(systemID, id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// Update updates a given BlockDevice
func (b *BlockDevice) Update(ctx context.Context, systemID string, id int, params *entity.BlockDeviceParams) (*entity.BlockDevice, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	blockDevice := new(entity.BlockDevice)
	err = b.client(systemID, id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// Delete deletes a given BlockDevice
func (b *BlockDevice) Delete(ctx context.Context, systemID string, id int) error {
	return b.client(systemID, id).Delete(ctx)
}

// AddTag adds a tag to a given BlockDevice
func (b *BlockDevice) AddTag(ctx context.Context, systemID string, id int, tag string) (*entity.BlockDevice, error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)

	blockDevice := new(entity.BlockDevice)
	err := b.client(systemID, id).Post(ctx, "add_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// RemoveTag removes a tag from a given BlockDevice
func (b *BlockDevice) RemoveTag(ctx context.Context, systemID string, id int, tag string) (*entity.BlockDevice, error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)

	blockDevice := new(entity.BlockDevice)
	err := b.client(systemID, id).Post(ctx, "remove_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// Format configures a BlockDevice to be formatted for a new file system
func (b *BlockDevice) Format(ctx context.Context, systemID string, id int, fsType string) (*entity.BlockDevice, error) {
	qsp := url.Values{}
	qsp.Set("fstype", fsType)

	blockDevice := new(entity.BlockDevice)
	err := b.client(systemID, id).Post(ctx, "format", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// Unformat removes the configured file system
func (b *BlockDevice) Unformat(ctx context.Context, systemID string, id int) (*entity.BlockDevice, error) {
	blockDevice := new(entity.BlockDevice)
	err := b.client(systemID, id).Post(ctx, "unformat", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// Mount sets the mount point and options of a given BlockDevice
func (b *BlockDevice) Mount(ctx context.Context, systemID string, id int, mountPoint string, mountOptions string) (*entity.BlockDevice, error) {
	qsp := url.Values{}
	qsp.Set("mount_point", mountPoint)
	qsp.Set("mount_options", mountOptions)

	blockDevice := new(entity.BlockDevice)
	err := b.client(systemID, id).Post(ctx, "mount", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// Unmount unsets the mount of a BlockDevice
func (b *BlockDevice) Unmount(ctx context.Context, systemID string, id int) (*entity.BlockDevice, error) {
	blockDevice := new(entity.BlockDevice)
	err := b.client(systemID, id).Post(ctx, "unmount", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// SetBootDisk configures the given BlockDevice as the boot disk
func (b *BlockDevice) SetBootDisk(ctx context.Context, systemID string, id int) error {
	return b.client(systemID, id).Post(ctx, "set_boot_disk", url.Values{}, func(data []byte) error { return nil })
}
