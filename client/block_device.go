package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// BlockDevice implements the api.BlockDevice interface
type BlockDevice struct {
	ApiClient ApiClient
}

func (b *BlockDevice) client(systemID string, id int) ApiClient {
	return b.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("blockdevices").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a given BlockDevice based on the given Node's system_id and BlockDevice's id
func (b *BlockDevice) Get(systemID string, id int) (blockDevice *entity.BlockDevice, err error) {
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// Update updates a given BlockDevice
func (b *BlockDevice) Update(systemID string, id int, params *entity.BlockDeviceParams) (blockDevice *entity.BlockDevice, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// Delete deletes a given BlockDevice
func (b *BlockDevice) Delete(systemID string, id int) error {
	return b.client(systemID, id).Delete()
}

// AddTag adds a tag to a given BlockDevice
func (b *BlockDevice) AddTag(systemID string, id int, tag string) (blockDevice *entity.BlockDevice, err error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("add_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// RemoveTag removes a tag from a given BlockDevice
func (b *BlockDevice) RemoveTag(systemID string, id int, tag string) (blockDevice *entity.BlockDevice, err error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("remove_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// Format configures a BlockDevice to be formatted for a new file system
func (b *BlockDevice) Format(systemID string, id int, fsType string) (blockDevice *entity.BlockDevice, err error) {
	qsp := url.Values{}
	qsp.Set("fstype", fsType)
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("format", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// Unformat removes the configured file system
func (b *BlockDevice) Unformat(systemID string, id int) (blockDevice *entity.BlockDevice, err error) {
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("unformat", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// Mount sets the mount point and options of a given BlockDevice
func (b *BlockDevice) Mount(systemID string, id int, mountPoint string, mountOptions string) (blockDevice *entity.BlockDevice, err error) {
	qsp := url.Values{}
	qsp.Set("mount_point", mountPoint)
	qsp.Set("mount_options", mountOptions)
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("mount", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// Unmount unsets the mount of a BlockDevice
func (b *BlockDevice) Unmount(systemID string, id int) (blockDevice *entity.BlockDevice, err error) {
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("unmount", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

// SetBootDisk configures the given BlockDevice as the boot disk
func (b *BlockDevice) SetBootDisk(systemID string, id int) error {
	return b.client(systemID, id).Post("set_boot_disk", url.Values{}, func(data []byte) error { return nil })
}
