package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ionutbalutoiu/gomaasclient/entity"
)

type VLANs struct {
	ApiClient ApiClient
}

func (v *VLANs) client(fabricID int) ApiClient {
	return v.ApiClient.GetSubObject("fabrics").GetSubObject(fmt.Sprintf("%v", fabricID)).GetSubObject("vlans")
}

func (v *VLANs) Get(fabricID int) (vlans []entity.VLAN, err error) {
	err = v.client(fabricID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &vlans)
	})
	return
}

func (v *VLANs) Create(fabricID int, params *entity.VLANParams) (vlan *entity.VLAN, err error) {
	vlan = new(entity.VLAN)
	err = v.client(fabricID).Post("", ToQSP(params), func(data []byte) error {
		return json.Unmarshal(data, vlan)
	})
	return
}
