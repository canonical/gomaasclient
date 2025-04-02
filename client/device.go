package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Device implements api.Device
type Device struct {
	APIClient APIClient
}

func (d *Device) client(systemID string) *APIClient {
	return d.APIClient.SubClient("devices").SubClient(fmt.Sprintf("%v", systemID))
}

// Get fetches a device with a given system_id
func (d *Device) Get(ctx context.Context, systemID string) (*entity.Device, error) {
	device := new(entity.Device)
	err := d.client(systemID).Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, device)
	})

	return device, err
}

// Update updates a given Device
func (d *Device) Update(ctx context.Context, systemID string, deviceParams *entity.DeviceUpdateParams) (*entity.Device, error) {
	qsp, err := query.Values(deviceParams)
	if err != nil {
		return nil, err
	}

	device := new(entity.Device)
	err = d.client(systemID).Put(ctx, qsp, func(data []byte) error {
		return json.Unmarshal(data, device)
	})

	return device, err
}

// Delete deletes a given Device
func (d *Device) Delete(ctx context.Context, systemID string) error {
	return d.client(systemID).Delete(ctx)
}

// SetWorkloadAnnotations add, modify or remove workload annotations for given Device
func (d *Device) SetWorkloadAnnotations(ctx context.Context, systemID string, params map[string]string) (*entity.Device, error) {
	qsp := url.Values{}
	for k, v := range params {
		qsp.Add(k, v)
	}

	device := new(entity.Device)
	err := d.client(systemID).Post(ctx, "set_workload_annotations", qsp, func(data []byte) error {
		return json.Unmarshal(data, &device)
	})

	return device, err
}
