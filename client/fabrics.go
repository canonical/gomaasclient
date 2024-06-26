//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Fabrics implements api.Fabrics
type Fabrics struct {
	APIClient APIClient
}

func (f *Fabrics) client() APIClient {
	return f.APIClient.GetSubObject("fabrics")
}

// Get fetches a list of Fabric objects
func (f *Fabrics) Get() ([]entity.Fabric, error) {
	fabrics := make([]entity.Fabric, 0)
	err := f.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &fabrics)
	})

	return fabrics, err
}

// Create creates a new Fabric object
func (f *Fabrics) Create(fabricParams *entity.FabricParams) (*entity.Fabric, error) {
	qsp, err := query.Values(fabricParams)
	if err != nil {
		return nil, err
	}

	fabric := new(entity.Fabric)
	err = f.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})

	return fabric, err
}
