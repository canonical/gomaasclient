package client

import (
	"encoding/json"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// IPAddresses implements api.IPAddresses
type IPAddresses struct {
	APIClient APIClient
}

func (i *IPAddresses) client() APIClient {
	return i.APIClient.GetSubObject("ipaddresses")
}

// Get fetches a specified set of IPAddresses
func (i *IPAddresses) Get(params *entity.IPAddressesParams) ([]entity.IPAddress, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	ipAddresses := make([]entity.IPAddress, 0)
	err = i.client().Get("", qsp, func(data []byte) error {
		return json.Unmarshal(data, &ipAddresses)
	})

	return ipAddresses, err
}

// Release releases a set of allocated IPAddresses
func (i *IPAddresses) Release(params *entity.IPAddressesParams) error {
	qsp, err := query.Values(params)
	if err != nil {
		return err
	}

	return i.client().Post("release", qsp, func(data []byte) error { return nil })
}

// Reserve reserves a set of IPAddresses
func (i *IPAddresses) Reserve(params *entity.IPAddressesParams) (*entity.IPAddress, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	ipAddress := new(entity.IPAddress)
	err = i.client().Post("reserve", qsp, func(data []byte) error {
		return json.Unmarshal(data, ipAddress)
	})

	return ipAddress, err
}
