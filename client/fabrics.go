package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type Fabrics struct {
	ApiClient ApiClient
}

func (f *Fabrics) client() ApiClient {
	return f.ApiClient.GetSubObject("fabrics")
}

func (f *Fabrics) Get() (fabrics []entity.Fabric, err error) {
	err = f.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &fabrics)
	})
	return
}

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
