package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// LicenseKey implements api.LicenseKey
type LicenseKey struct {
	APIClient APIClient
}

func (f *LicenseKey) client(osystem string, distroSeries string) APIClient {
	return f.APIClient.
		GetSubObject("license-key").
		GetSubObject(osystem).
		GetSubObject(distroSeries)
}

// Get fetches a LicenseKey object
func (f *LicenseKey) Get(osystem string, distroSeries string) (*entity.LicenseKey, error) {
	licenseKey := new(entity.LicenseKey)
	err := f.client(osystem, distroSeries).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, licenseKey)
	})

	return licenseKey, err
}

// Update updates a given LicenseKey
func (l *LicenseKey) Update(osystem string, distroSeries string, params *entity.LicenseKeyParams) (*entity.LicenseKey, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	licenseKey := new(entity.LicenseKey)
	err = l.client(osystem, distroSeries).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, licenseKey)
	})

	return licenseKey, err
}

// Delete deletes a given LicenseKey
func (l *LicenseKey) Delete(osystem string, distroSeries string) error {
	return l.client(osystem, distroSeries).Delete()
}
