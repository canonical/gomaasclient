package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// ReservedIPs is an interface for listing and creating
// ReservedIP records
type ReservedIPs interface {
	Get() ([]entity.ReservedIP, error)
	Create(params *entity.ReservedIPCreateParams) (*entity.ReservedIP, error)
}
