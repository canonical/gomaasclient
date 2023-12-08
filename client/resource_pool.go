package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// ResourcePool implements api.ResourcePool
type ResourcePool struct {
	APIClient APIClient
}

func (r *ResourcePool) client(id int) APIClient {
	return r.APIClient.GetSubObject(fmt.Sprintf("resourcepool/%v", id))
}

// Get fetches a given ResourcePool
func (r *ResourcePool) Get(id int) (resourcePool *entity.ResourcePool, err error) {
	resourcePool = new(entity.ResourcePool)
	err = r.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})

	return
}

// Update updates a given ResourcePool
func (r *ResourcePool) Update(id int, params *entity.ResourcePoolParams) (resourcePool *entity.ResourcePool, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}

	resourcePool = new(entity.ResourcePool)
	err = r.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})

	return
}

// Delete deletes a given ResourcePool
func (r *ResourcePool) Delete(id int) error {
	return r.client(id).Delete()
}
