package api

import (
	"github.com/canonical/gomaasclient/entity"
)

// Zone is an interface defining API behaviour for zones
type Zone interface {
	Get(name string) (*entity.Zone, error)
	Update(name string, params *entity.ZoneParams) (*entity.Zone, error)
	Delete(name string) error
}
