package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

type ResourcePool struct {
	ApiClient ApiClient
}

func (r *ResourcePool) client(id int) ApiClient {
	return r.ApiClient.GetSubObject(fmt.Sprintf("resourcepool/%v", id))
}

func (r *ResourcePool) Get(id int) (resourcePool *entity.ResourcePool, err error) {
	resourcePool = new(entity.ResourcePool)
	err = r.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, resourcePool)
	})
	return
}

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

func (r *ResourcePool) Delete(id int) error {
	return r.client(id).Delete()
}
