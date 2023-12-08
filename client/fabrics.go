//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Fabrics implements api.Fabrics
type Fabrics struct {
	APIClient APIClient
}

func (f *Fabrics) client() APIClient {
	return f.APIClient.GetSubObject("fabrics")
}

// Get fetches a list of Fabric objects
func (f *Fabrics) Get() (fabrics []entity.Fabric, err error) {
	err = f.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &fabrics)
	})

	return
}

// Create creates a new Fabric object
func (f *Fabrics) Create(fabricParams *entity.FabricParams) (fabric *entity.Fabric, err error) {
	qsp, err := query.Values(fabricParams)
	if err != nil {
		return
	}

	fabric = new(entity.Fabric)
	err = f.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})

	return
}
