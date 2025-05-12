package api

import "github.com/canonical/gomaasclient/entity"

// LicenseKeys is an interface for listing and creating LicenseKey objects
type LicenseKeys interface {
	Get() ([]entity.LicenseKey, error)
	Create(params *entity.LicenseKeyParams) (*entity.LicenseKey, error)
}
