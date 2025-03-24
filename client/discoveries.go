package client

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// Discoveries implements api.Discoveries
type Discoveries struct {
	APIClient APIClient
}

func (d *Discoveries) client() *APIClient {
	return d.APIClient.SubClient("discovery")
}

// Get complete discovery list
func (d *Discoveries) Get(ctx context.Context) ([]entity.Discovery, error) {
	deviceDiscovery := make([]entity.Discovery, 0)
	err := d.client().Get(ctx, "", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &deviceDiscovery)
	})

	return deviceDiscovery, err
}

// GetByUnknownIP get discovery list with unknown IP
func (d *Discoveries) GetByUnknownIP(ctx context.Context) ([]entity.Discovery, error) {
	deviceDiscovery := make([]entity.Discovery, 0)
	err := d.client().Get(ctx, "by_unknown_ip", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &deviceDiscovery)
	})

	return deviceDiscovery, err
}

// GetByUnknownMAC get discovery list with unknown MAC
func (d *Discoveries) GetByUnknownMAC(ctx context.Context) ([]entity.Discovery, error) {
	deviceDiscovery := make([]entity.Discovery, 0)
	err := d.client().Get(ctx, "by_unknown_mac", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &deviceDiscovery)
	})

	return deviceDiscovery, err
}

// GetByUnknownIPAndMAC get discovery list with unknown IP and MAC
func (d *Discoveries) GetByUnknownIPAndMAC(ctx context.Context) ([]entity.Discovery, error) {
	deviceDiscovery := make([]entity.Discovery, 0)
	err := d.client().Get(ctx, "by_unknown_ip_and_mac", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &deviceDiscovery)
	})

	return deviceDiscovery, err
}

// Clear deletes all discovered neighbours and/or mDNS entries.
func (d *Discoveries) Clear(ctx context.Context, params *entity.DiscoveryClearParams) error {
	qsp, err := query.Values(params)
	if err != nil {
		return err
	}

	return d.client().Post(ctx, "clear", qsp, func(data []byte) error { return nil })
}

// ClearByMACAndIP delete discoveries that match a MAC and IP
func (d *Discoveries) ClearByMACAndIP(ctx context.Context, mac, ip string) error {
	qsp := url.Values{}
	qsp.Set("mac", mac)
	qsp.Set("ip", ip)

	return d.client().Post(ctx, "clear_by_mac_and_ip", qsp, func(data []byte) error { return nil })
}
