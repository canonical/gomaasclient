package client

import (
	"context"
	"encoding/json"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// IPAddresses implements api.IPAddresses
type IPAddresses struct {
	APIClient APIClient
}

func (i *IPAddresses) client() *APIClient {
	return i.APIClient.SubClient("ipaddresses")
}

// Get fetches a specified set of IPAddresses
func (i *IPAddresses) Get(ctx context.Context, params *entity.IPAddressesParams) ([]entity.IPAddress, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	ipAddresses := make([]entity.IPAddress, 0)
	err = i.client().Get(ctx, "", qsp, func(data []byte) error {
		return json.Unmarshal(data, &ipAddresses)
	})

	return ipAddresses, err
}

// Release releases a set of allocated IPAddresses
func (i *IPAddresses) Release(ctx context.Context, params *entity.IPAddressesParams) error {
	qsp, err := query.Values(params)
	if err != nil {
		return err
	}

	return i.client().Post(ctx, "release", qsp, func(data []byte) error { return nil })
}

// Reserve reserves a set of IPAddresses
func (i *IPAddresses) Reserve(ctx context.Context, params *entity.IPAddressesParams) (*entity.IPAddress, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	ipAddress := new(entity.IPAddress)
	err = i.client().Post(ctx, "reserve", qsp, func(data []byte) error {
		return json.Unmarshal(data, ipAddress)
	})

	return ipAddress, err
}
