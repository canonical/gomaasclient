package client

import (
	"context"
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

func (v *VolumeGroup) client(systemID string, id int) *APIClient {
	return v.APIClient.SubClient("nodes").SubClient(systemID).SubClient("volume-groups").SubClient(fmt.Sprintf("%v", id))
}

// Get fetches a given VolumeGroup based on the given Node's system_id and VolumeGroup's ID
func (v *VolumeGroup) Get(ctx context.Context, systemID string, id int) (*entity.VolumeGroup, error) {
	volumeGroup := new(entity.VolumeGroup)
	err := v.client(systemID, id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, volumeGroup)
	})

	return volumeGroup, err
}

// Update updates a given VolumeGroup
func (v *VolumeGroup) Update(ctx context.Context, systemID string, id int, params *entity.VolumeGroupUpdateParams) (*entity.VolumeGroup, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	volumeGroup := new(entity.VolumeGroup)
	err = v.client(systemID, id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, volumeGroup)
	})

	return volumeGroup, err
}

// Delete deletes a given VolumeGroup
func (v *VolumeGroup) Delete(ctx context.Context, systemID string, id int) error {
	return v.client(systemID, id).Delete(ctx)
}

// CreateLogicalVolume creates a new LogicalVolume for a given system_id and VolumeGroup's ID
func (v *VolumeGroup) CreateLogicalVolume(ctx context.Context, systemID string, id int, params *entity.LogicalVolumeParams) (*entity.BlockDevice, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	blockDevice := new(entity.BlockDevice)
	err = v.client(systemID, id).Post(ctx, "create_logical_volume", qsp, func(data []byte) error {
		return json.Unmarshal(data, blockDevice)
	})

	return blockDevice, err
}

// DeleteLogicalVolume deletes a given LogicalVolume
func (v *VolumeGroup) DeleteLogicalVolume(ctx context.Context, systemID string, id int, lvID int) error {
	qsp := url.Values{}
	qsp.Set("id", strconv.Itoa(lvID))

	err := v.client(systemID, id).Post(ctx, "delete_logical_volume", qsp, func(data []byte) error {
		return nil
	})

	return err
}
