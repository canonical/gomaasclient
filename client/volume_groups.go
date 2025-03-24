//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// VolumeGroups implements the api.VolumeGroups interface
type VolumeGroups struct {
	APIClient APIClient
}

func (v *VolumeGroups) client(systemID string) *APIClient {
	return v.APIClient.SubClient("nodes").SubClient(systemID).SubClient("volume-groups")
}

// Get fetches a list of VolumeGroups for a given system_id
func (v *VolumeGroups) Get(ctx context.Context, systemID string) ([]entity.VolumeGroup, error) {
	volumeGroups := make([]entity.VolumeGroup, 0)
	err := v.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &volumeGroups)
	})

	return volumeGroups, err
}

// Create creates a new VolumeGroup for a given system_id
func (v *VolumeGroups) Create(ctx context.Context, systemID string, params *entity.VolumeGroupCreateParams) (*entity.VolumeGroup, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	volumeGroup := new(entity.VolumeGroup)
	err = v.client(systemID).Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, volumeGroup)
	})

	return volumeGroup, err
}
