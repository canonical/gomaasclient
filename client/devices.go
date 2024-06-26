//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Devices implements api.Devices
type Devices struct {
	APIClient APIClient
}

func (d *Devices) client() APIClient {
	return d.APIClient.GetSubObject("devices")
}

// Get fetches a list of Devices
func (d *Devices) Get() ([]entity.Device, error) {
	devices := make([]entity.Device, 0)
	err := d.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &devices)
	})

	return devices, err
}

// Create creates a new Device
func (d *Devices) Create(deviceParams *entity.DeviceCreateParams) (*entity.Device, error) {
	qsp, err := query.Values(deviceParams)
	if err != nil {
		return nil, err
	}

	device := new(entity.Device)
	err = d.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, device)
	})

	return device, err
}
