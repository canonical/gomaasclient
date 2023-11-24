package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

type ResourcePools struct {
	ApiClient ApiClient
}

func (r *ResourcePools) client() ApiClient {
	return r.ApiClient.GetSubObject("resourcepools")
}

func (r *ResourcePools) Get() (resourcePools []entity.ResourcePool, err error) {
	err = r.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &resourcePools)
	})
	return
}

func (r *ResourcePools) Create(params *entity.ResourcePoolParams) (resourcePool *entity.ResourcePool, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	resourcePool = new(entity.ResourcePool)
	err = r.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})
	return
}
