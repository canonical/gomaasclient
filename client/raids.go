//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/canonical/gomaasclient/entity"
)

// RAIDs contains functionality for manipulating the RAIDs entity.
type RAIDs struct {
	APIClient APIClient
}

func (r *RAIDs) client(systemID string) *APIClient {
	return r.APIClient.SubClient("nodes").SubClient(systemID).SubClient("raids")
}

// Get RAIDs of a machine.
func (r *RAIDs) Get(ctx context.Context, systemID string) ([]entity.RAID, error) {
	raids := make([]entity.RAID, 0)
	err := r.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &raids)
	})

	return raids, err
}

// Create a RAID of a machine.
func (r *RAIDs) Create(ctx context.Context, systemID string, params *entity.RAIDCreateParams) (*entity.RAID, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	raid := new(entity.RAID)
	err = r.client(systemID).Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})

	return raid, err
}
