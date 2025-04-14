package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// ResourcePool implements api.ResourcePool
type ResourcePool struct {
	APIClient APIClient
}

func (r *ResourcePool) client(id int) *APIClient {
	return r.APIClient.SubClient(fmt.Sprintf("resourcepool/%v", id))
}

func (r *ResourcePool) clientByName(name string) *APIClient {
	return r.APIClient.SubClient(fmt.Sprintf("resourcepool/%v", name))
}

// GetByName fetches a given ResourcePool
func (r *ResourcePool) GetByName(ctx context.Context, name string) (*entity.ResourcePool, error) {
	resourcePool := new(entity.ResourcePool)
	err := r.clientByName(name).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})

	return resourcePool, err
}

// UpdateByName updates a given ResourcePool
func (r *ResourcePool) UpdateByName(ctx context.Context, name string, params *entity.ResourcePoolParams) (*entity.ResourcePool, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	resourcePool := new(entity.ResourcePool)
	err = r.clientByName(name).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})

	return resourcePool, err
}

// DeleteByName deletes a given ResourcePool
func (r *ResourcePool) DeleteByName(ctx context.Context, name string) error {
	return r.clientByName(name).Delete(ctx)
}

// Get fetches a given ResourcePool
func (r *ResourcePool) Get(ctx context.Context, id int) (*entity.ResourcePool, error) {
	resourcePool := new(entity.ResourcePool)
	err := r.client(id).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})

	return resourcePool, err
}

// Update updates a given ResourcePool
func (r *ResourcePool) Update(ctx context.Context, id int, params *entity.ResourcePoolParams) (*entity.ResourcePool, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	resourcePool := new(entity.ResourcePool)
	err = r.client(id).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})

	return resourcePool, err
}

// Delete deletes a given ResourcePool
func (r *ResourcePool) Delete(ctx context.Context, id int) error {
	return r.client(id).Delete(ctx)
}
