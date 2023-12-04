package client

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// IPAddresses implements api.IPAddresses
type IPAddresses struct {
	ApiClient ApiClient
}

func (i *IPAddresses) client() ApiClient {
	return i.ApiClient.GetSubObject("ipaddresses")
}

// Get fetches a specified set of IPAddresses
func (i *IPAddresses) Get(params *entity.IPAddressesParams) (ipAddresses []entity.IPAddress, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	err = i.client().Get("", qsp, func(data []byte) error {
		return json.Unmarshal(data, &ipAddresses)
	})
	return
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
func (i *IPAddresses) Reserve(params *entity.IPAddressesParams) (ipAddress *entity.IPAddress, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	ipAddress = new(entity.IPAddress)
	err = i.client().Post("reserve", qsp, func(data []byte) error {
		return json.Unmarshal(data, ipAddress)
	})
	return
}
