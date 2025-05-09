package api

import "github.com/canonical/gomaasclient/entity"

// LicenseKey is an interface defining API behavior for
// LicenseKey objects
type LicenseKey interface {
	Get(osystem string, distroSeries string) (*entity.LicenseKey, error)
	Update(osystem string, distroSeries string, params *entity.LicenseKeyParams) (*entity.LicenseKey, error)
	Delete(osystem string, distroSeries string) error
}
