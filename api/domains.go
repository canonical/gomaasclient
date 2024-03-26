package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Domains is an interface for listing and creaing
// Domain records
type Domains interface {
	Get() ([]entity.Domain, error)
	Create(params *entity.DomainParams) (*entity.Domain, error)
	SetSerial(serial int) error
}
