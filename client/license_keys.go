//nolint:dupl // disable dupl check on client for now
package client

import (
	"encoding/json"
	"net/url"

	"github.com/canonical/gomaasclient/entity"
	"github.com/google/go-querystring/query"
)

// LicenseKeys implements api.LicenseKeys
type LicenseKeys struct {
	APIClient APIClient
}

func (l *LicenseKeys) client() APIClient {
	return l.APIClient.GetSubObject("license-keys")
}

// Get fetches a list of LicenseKey objects
func (l *LicenseKeys) Get() ([]entity.LicenseKey, error) {
	licenseKeys := make([]entity.LicenseKey, 0)
	err := l.client().Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &licenseKeys)
	})

	return licenseKeys, err
}

// Create creates a new LicenseKey object
func (l *LicenseKeys) Create(params *entity.LicenseKeyParams) (*entity.LicenseKey, error) {
	qsp, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	licenseKey := new(entity.LicenseKey)
	err = l.client().Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, licenseKey)
	})

	return licenseKey, err
}
