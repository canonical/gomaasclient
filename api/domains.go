package api

import (
	"github.com/maas/gomaasclient/entity"
)

// Domains is an interface for listing and creaing
// Domain records
type Domains interface {
	Get() ([]entity.Domain, error)
	Create(params *entity.DomainParams) (*entity.Domain, error)
	SetSerial(serial int) error
}
