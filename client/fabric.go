package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Fabric struct {
	ApiClient ApiClient
}

func (f *Fabric) client(id int) ApiClient {
	return f.ApiClient.GetSubObject("fabrics").GetSubObject(fmt.Sprintf("%v", id))
}

func (f *Fabric) Get(id int) (fabric *entity.Fabric, err error) {
	fabric = new(entity.Fabric)
	err = f.client(id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})
	return
}

func (f *Fabric) Update(id int, fabricParams *entity.FabricParams) (fabric *entity.Fabric, err error) {
	fabric = new(entity.Fabric)
	err = f.client(id).Put(ToQSP(fabricParams), func(data []byte) error {
		return json.Unmarshal(data, fabric)
	})
	return
}

func (f *Fabric) Delete(id int) error {
	return f.client(id).Delete()
}
