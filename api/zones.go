package api

import (
	"github.com/maas/gomaasclient/entity"
)

// Zones is an interface for listing and creating Zone objects
type Zones interface {
	Get() ([]entity.Zone, error)
	Create(params *entity.ZoneParams) (*entity.Zone, error)
}
