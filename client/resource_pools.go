//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// ResourcePools implements api.ResourcePools
type ResourcePools struct {
	APIClient APIClient
}

func (r *ResourcePools) client() APIClient {
	return r.APIClient.GetSubObject("resourcepools")
}

// Get fetches a list of ResourcePool objects
func (r *ResourcePools) Get() ([]entity.ResourcePool, error) {
	resourcePools := make([]entity.ResourcePool, 0)
	err := r.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &resourcePools)
	})

	return resourcePools, err
}

// Create creates a new ResourcePool
func (r *ResourcePools) Create(params *entity.ResourcePoolParams) (*entity.ResourcePool, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	resourcePool := new(entity.ResourcePool)
	err = r.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})

	return resourcePool, err
}
