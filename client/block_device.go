package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type BlockDevice struct {
	ApiClient ApiClient
}

func (b *BlockDevice) client(systemID string, id int) ApiClient {
	return b.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("blockdevices").GetSubObject(fmt.Sprintf("%v", id))
}

func (b *BlockDevice) Get(systemID string, id int) (blockDevice *entity.BlockDevice, err error) {
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

func (b *BlockDevice) Update(systemID string, id int, params *entity.BlockDeviceParams) (blockDevice *entity.BlockDevice, err error) {
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Put(ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

func (b *BlockDevice) Delete(systemID string, id int) error {
	return b.client(systemID, id).Delete()
}

func (b *BlockDevice) AddTag(systemID string, id int, tag string) (blockDevice *entity.BlockDevice, err error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("add_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

func (b *BlockDevice) RemoveTag(systemID string, id int, tag string) (blockDevice *entity.BlockDevice, err error) {
	qsp := url.Values{}
	qsp.Set("tag", tag)
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("remove_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

func (b *BlockDevice) Format(systemID string, id int, fsType string) (blockDevice *entity.BlockDevice, err error) {
	qsp := url.Values{}
	qsp.Set("fstype", fsType)
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("format", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

func (b *BlockDevice) Unformat(systemID string, id int) (blockDevice *entity.BlockDevice, err error) {
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("unformat", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

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

func (b *BlockDevice) Unmount(systemID string, id int) (blockDevice *entity.BlockDevice, err error) {
	blockDevice = new(entity.BlockDevice)
	err = b.client(systemID, id).Post("unmount", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})
	return
}

func (b *BlockDevice) SetBootDisk(systemID string, id int) error {
	return b.client(systemID, id).Post("set_boot_disk", url.Values{}, func(data []byte) error { return nil })
}
