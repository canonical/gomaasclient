package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// VolumeGroup implements the api.VolumeGroup interface
type VolumeGroup struct {
	APIClient APIClient
}

func (v *VolumeGroup) client(systemID string, id int) APIClient {
	return v.APIClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("volume-groups").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetches a given VolumeGroup based on the given Node's system_id and VolumeGroup's ID
func (v *VolumeGroup) Get(systemID string, id int) (*entity.VolumeGroup, error) {
	volumeGroup := new(entity.VolumeGroup)
	err := v.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, volumeGroup)
	})

	return volumeGroup, err
}

// Update updates a given VolumeGroup
func (v *VolumeGroup) Update(systemID string, id int, params *entity.VolumeGroupParams) (*entity.VolumeGroup, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	volumeGroup := new(entity.VolumeGroup)
	err = v.client(systemID, id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, volumeGroup)
	})

	return volumeGroup, err
}

// Delete deletes a given VolumeGroup
func (v *VolumeGroup) Delete(systemID string, id int) error {
	return v.client(systemID, id).Delete()
}

// CreateLogicalVolume creates a new LogicalVolume for a given system_id and VolumeGroup's ID
func (v *VolumeGroup) CreateLogicalVolume(systemID string, id int, params *entity.LogicalVolumeParams) (*entity.LogicalVolume, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	logicalVolume := new(entity.LogicalVolume)
	err = v.client(systemID, id).Post("create_logical_volume", qsp, func(data []byte) error {
		return json.Unmarshal(data, logicalVolume)
	})

	return logicalVolume, err
}

// DeleteLogicalVolume deletes a given LogicalVolume
func (v *VolumeGroup) DeleteLogicalVolume(systemID string, id int, lvID int) error {
	qsp := url.Values{}
	qsp.Set("id", strconv.Itoa(lvID))

	err := v.client(systemID, id).Post("delete_logical_volume", qsp, func(data []byte) error {
		return nil
	})

	return err
}
