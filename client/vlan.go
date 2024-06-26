package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// VLAN implements api.VLAN
type VLAN struct {
	APIClient APIClient
}

func (v *VLAN) client(fabricID int, vid int) APIClient {
	return v.APIClient.
		GetSubObject("fabrics").
		GetSubObject(fmt.Sprintf("%v", fabricID)).
		GetSubObject("vlans").
		GetSubObject(fmt.Sprintf("%v", vid))
}

// Get fetches a VLAN by fabric id and VLAN id
func (v *VLAN) Get(fabricID int, vid int) (*entity.VLAN, error) {
	vlan := new(entity.VLAN)
	err := v.client(fabricID, vid).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, vlan)
	})

	return vlan, err
}

// Update updates a given VLAN
func (v *VLAN) Update(fabricID int, vid int, params *entity.VLANParams) (*entity.VLAN, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	vlan := new(entity.VLAN)
	err = v.client(fabricID, vid).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, vlan)
	})

	return vlan, err
}

// Delete deletes a given VLAN
func (v *VLAN) Delete(fabricID int, vid int) error {
	return v.client(fabricID, vid).Delete()
}
