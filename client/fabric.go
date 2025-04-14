//nolint:dupl // disable dupl check on client for now
package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Fabric implments api.Fabric
type Fabric struct {
	APIClient APIClient
}

func (f *Fabric) client(id int) *APIClient {
	return f.APIClient.SubClient("fabrics").SubClient(fmt.Sprintf("%v", id))
}

// Get fetchs a Fabric object
func (f *Fabric) Get(ctx context.Context, id int) (*entity.Fabric, error) {
	fabric := new(entity.Fabric)
	err := f.client(id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})

	return fabric, err
}

// Update updates a given Fabric
func (f *Fabric) Update(ctx context.Context, id int, fabricParams *entity.FabricParams) (*entity.Fabric, error) {
	qsp, err := query.Values(fabricParams)
	if err != nil {
		return nil, err
	}

	fabric := new(entity.Fabric)
	err = f.client(id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})

	return fabric, err
}

// Delete deletes a given Fabric
func (f *Fabric) Delete(ctx context.Context, id int) error {
	return f.client(id).Delete(ctx)
}
