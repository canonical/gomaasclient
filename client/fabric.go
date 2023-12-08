package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Fabric implments api.Fabric
type Fabric struct {
	ApiClient ApiClient
}

func (f *Fabric) client(id int) ApiClient {
	return f.ApiClient.GetSubObject("fabrics").GetSubObject(fmt.Sprintf("%v", id))
}

// Get fetchs a Fabric object
func (f *Fabric) Get(id int) (fabric *entity.Fabric, err error) {
	fabric = new(entity.Fabric)
	err = f.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})

	return
}

// Update updates a given Fabric
func (f *Fabric) Update(id int, fabricParams *entity.FabricParams) (fabric *entity.Fabric, err error) {
	qsp, err := query.Values(fabricParams)
	if err != nil {
		return
	}

	fabric = new(entity.Fabric)
	err = f.client(id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})

	return
}

// Delete deletes a given Fabric
func (f *Fabric) Delete(id int) error {
	return f.client(id).Delete()
}
