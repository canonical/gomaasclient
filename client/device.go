package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Device implements api.Device
type Device struct {
	APIClient APIClient
}

func (d *Device) client(systemID string) APIClient {
	return d.APIClient.GetSubObject("devices").GetSubObject(fmt.Sprintf("%v", systemID))
}

// Get fetches a device with a given system_id
func (d *Device) Get(systemID string) (*entity.Device, error) {
	device := new(entity.Device)
	err := d.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, device)
	})

	return device, err
}

// Update updates a given Device
func (d *Device) Update(systemID string, deviceParams *entity.DeviceUpdateParams) (*entity.Device, error) {
	qsp, err := query.Values(deviceParams)
	if err != nil {
		return nil, err
	}

	device := new(entity.Device)
	err = d.client(systemID).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, device)
	})

	return device, err
}

// Delete deletes a given Device
func (d *Device) Delete(systemID string) error {
	return d.client(systemID).Delete()
}
