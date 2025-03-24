//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Fabrics implements api.Fabrics
type Fabrics struct {
	APIClient APIClient
}

func (f *Fabrics) client() *APIClient {
	return f.APIClient.SubClient("fabrics")
}

// Get fetches a list of Fabric objects
func (f *Fabrics) Get(ctx context.Context) ([]entity.Fabric, error) {
	fabrics := make([]entity.Fabric, 0)
	err := f.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &fabrics)
	})

	return fabrics, err
}

// Create creates a new Fabric object
func (f *Fabrics) Create(ctx context.Context, fabricParams *entity.FabricParams) (*entity.Fabric, error) {
	qsp, err := query.Values(fabricParams)
	if err != nil {
		return nil, err
	}

	fabric := new(entity.Fabric)
	err = f.client().Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})

	return fabric, err
}
