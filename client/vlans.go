package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// VLANs implements api.VLANs
type VLANs struct {
	APIClient APIClient
}

func (v *VLANs) client(fabricID int) *APIClient {
	return v.APIClient.SubClient("fabrics").SubClient(fmt.Sprintf("%v", fabricID)).SubClient("vlans")
}

// Get fetches a list of VLAN objects
func (v *VLANs) Get(ctx context.Context, fabricID int) ([]entity.VLAN, error) {
	vlans := make([]entity.VLAN, 0)
	err := v.client(fabricID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &vlans)
	})

	return vlans, err
}

// Create creates a new VLAN
func (v *VLANs) Create(ctx context.Context, fabricID int, params *entity.VLANParams) (*entity.VLAN, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	vlan := new(entity.VLAN)
	err = v.client(fabricID).Post(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, vlan)
	})

	return vlan, err
}
