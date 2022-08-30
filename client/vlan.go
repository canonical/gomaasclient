package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

type VLAN struct {
	ApiClient ApiClient
}

func (v *VLAN) client(fabricID int, vid int) ApiClient {
	return v.ApiClient.
		GetSubObject("fabrics").
		GetSubObject(fmt.Sprintf("%v", fabricID)).
		GetSubObject("vlans").
		GetSubObject(fmt.Sprintf("%v", vid))
}

func (v *VLAN) Get(fabricID int, vid int) (vlan *entity.VLAN, err error) {
	vlan = new(entity.VLAN)
	err = v.client(fabricID, vid).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, vlan)
	})
	return
}

func (v *VLAN) Update(fabricID int, vid int, params *entity.VLANParams) (vlan *entity.VLAN, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	vlan = new(entity.VLAN)
	err = v.client(fabricID, vid).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, vlan)
	})
	return
}

func (v *VLAN) Delete(fabricID int, vid int) error {
	return v.client(fabricID, vid).Delete()
}
