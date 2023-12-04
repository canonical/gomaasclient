package client

import (
	"encoding/json"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Devices implements api.Devices
type Devices struct {
	ApiClient ApiClient
}

func (d *Devices) client() ApiClient {
	return d.ApiClient.GetSubObject("devices")
}

// Get fetches a list of Devices
func (d *Devices) Get() (devices []entity.Device, err error) {
	err = d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &devices)
	})
	return
}

// Create creates a new Device
func (d *Devices) Create(deviceParams *entity.DeviceCreateParams) (device *entity.Device, err error) {
	qsp, err := query.Values(deviceParams)
	if err != nil {
		return
	}
	device = new(entity.Device)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, device)
	})
	return
}
