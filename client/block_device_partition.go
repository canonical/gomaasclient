package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	gomaasapi "github.com/juju/gomaasapi/v2"
	"github.com/maas/gomaasclient/entity"
)

// BlockDevicePartition implements the api.BlockDevicePartition
// interface
type BlockDevicePartition struct {
	ApiClient ApiClient
}

func (p *BlockDevicePartition) client(systemID string, blockDeviceID int, id int) (*ApiClient, error) {
	uri := p.ApiClient.URI()
	newURL := url.URL{Path: fmt.Sprintf("nodes/%s/blockdevices/%v/partition/%v", systemID, blockDeviceID, id)}
	resUrl := uri.ResolveReference(&newURL)
	input := map[string]interface{}{"resource_uri": resUrl.String()}
	jsonObj, err := gomaasapi.JSONObjectFromStruct(p.ApiClient.AuthClient, input)
	if err != nil {
		return nil, err
	}
	maasObj, err := jsonObj.GetMAASObject()
	if err != nil {
		return nil, err
	}
	return &ApiClient{p.ApiClient.AuthClient, &maasObj}, nil
}

// Get fetches a given BlockDevicePartion for a given system_id, BlockDevice id and partition id
func (p *BlockDevicePartition) Get(systemID string, blockDeviceID int, id int) (partition *entity.BlockDevicePartition, err error) {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return
	}
	partition = new(entity.BlockDevicePartition)
	err = client.Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}

// Delete deletes a given BlockDevicePartition
func (p *BlockDevicePartition) Delete(systemID string, blockDeviceID int, id int) error {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return err
	}
	return client.Delete()
}

// AddTag adds a tag to a given BlockDevicePartition
func (p *BlockDevicePartition) AddTag(systemID string, blockDeviceID int, id int, tag string) (partition *entity.BlockDevicePartition, err error) {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return
	}
	qsp := url.Values{}
	qsp.Set("tag", tag)
	partition = new(entity.BlockDevicePartition)
	err = client.Post("add_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}

// RemoveTag removes a tag from a given BlockDevicePartition
func (p *BlockDevicePartition) RemoveTag(systemID string, blockDeviceID int, id int, tag string) (partition *entity.BlockDevicePartition, err error) {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return
	}
	qsp := url.Values{}
	qsp.Set("tag", tag)
	partition = new(entity.BlockDevicePartition)
	err = client.Post("remove_tag", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}

// Format sets the BlockDevicePartition to be formatted with a given file system
func (p *BlockDevicePartition) Format(systemID string, blockDeviceID int, id int, fsType string, label string) (partition *entity.BlockDevicePartition, err error) {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return
	}
	qsp := url.Values{}
	qsp.Set("fstype", fsType)
	qsp.Set("label", label)
	partition = new(entity.BlockDevicePartition)
	err = client.Post("format", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}

// Unformat unsets a file system for a given BlockDevicePartition
func (p *BlockDevicePartition) Unformat(systemID string, blockDeviceID int, id int) (partition *entity.BlockDevicePartition, err error) {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return
	}
	partition = new(entity.BlockDevicePartition)
	err = client.Post("unformat", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}

// Mount sets a mount point and options for a given BlockDevicePartition
func (p *BlockDevicePartition) Mount(systemID string, blockDeviceID int, id int, mountPoint string, mountOptions string) (partition *entity.BlockDevicePartition, err error) {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return
	}
	qsp := url.Values{}
	qsp.Set("mount_point", mountPoint)
	qsp.Set("mount_options", mountOptions)
	partition = new(entity.BlockDevicePartition)
	err = client.Post("mount", qsp, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}

// Unmount unsets a mount point for a given BlockDevicePartition
func (p *BlockDevicePartition) Unmount(systemID string, blockDeviceID int, id int) (partition *entity.BlockDevicePartition, err error) {
	client, err := p.client(systemID, blockDeviceID, id)
	if err != nil {
		return
	}
	partition = new(entity.BlockDevicePartition)
	err = client.Post("unmount", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, partition)
	})
	return
}
